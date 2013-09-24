package swcg

import "fmt"
import "strconv"
import "sort"

type CardMap        	map[int]*Card
type ObjectiveSetDB 	[6]*Card
type SetMap         	map[int]*ObjectiveSetDB
type TypeMap        	map[CardType][]*Card
type KeywordMap     	map[CardKeywordType][]*Card
type TraitMap       	map[CardTraitType][]*Card
type PlayAreaSynergyMap []*Card

type Data interface{
	Print() string
	IntValue() int
}
type IntData struct {V int}
func (d IntData) Print() string {return strconv.Itoa(d.V)}
func (d IntData) IntValue() int {return d.V}

type StrData struct {V string}
func (d StrData) Print() string {return d.V}
func (d StrData) IntValue() int {return int([]byte(d.V)[0])}


type DataRow []Data
type Header struct {
	Name string
}
type DataCollection     struct {
	header        	[]Header
	rows          	[]DataRow
	sortIndices     []int
	lessF           func(i,j int) bool
}
func CreateDataCollection(h ...string) *DataCollection{
	d := new(DataCollection)
	d.header = make([]Header, len(h))
	for i, head := range h {
		d.header[i] = Header{Name: head}
	}
	d.lessF = func(i,j int) bool {return false} // default
	return d
}
func (d *DataCollection) Len() int  { return len(d.rows) }
func (d *DataCollection) Swap(i, j int) {
	d.rows[i], d.rows[j] = d.rows[j], d.rows[i]
}
func (d *DataCollection) Less(i, j int) bool {
	for _, index := range d.sortIndices {
		iVal, jVal := d.rows[i][index].IntValue(), d.rows[j][index].IntValue()
		if iVal != jVal {
			return d.lessF(iVal, jVal)
		}
	}
	return false
}
func (d *DataCollection) Sort(less func(i,j int) bool, dataIndices ...int) {
	if len(dataIndices) < 1 { panic("Need at least one data index to sort the data collection...") }

	d.lessF = less
	d.sortIndices = dataIndices
	sort.Sort(d)
}
func (d *DataCollection) AddRow(rawrow ...interface{}) {
	if len(rawrow) != len(d.header) {
		panic(fmt.Sprintf("Can't create row, different size from header (row: %v, header: %v)", rawrow, d.header))
	}
	row := make([]Data, len(rawrow))
	for i, rdata := range rawrow {
		switch typeData := rdata.(type) {
		case int:    row[i] = IntData{V: typeData}
		case string: row[i] = StrData{V: typeData}
		default:
			panic(fmt.Sprintf("Unkown Data type when building row (data: %v, row: %v)", rdata, rawrow))
		}
	}
	d.rows = append(d.rows, row)
}
func (d *DataCollection) FilterRow(predicate func(*DataRow) bool) {
	filteredRows := make([]DataRow, 0)
	for _, r := range d.rows {
		if predicate(&r) {
			filteredRows = append(filteredRows, r)
		}
	}
	d.rows = filteredRows
	
}
func (d *DataCollection) Print() string {
	out := ""
	for i, h := range d.header {
		if i < len(d.header)-1 {
			out += tabifyName(h.Name)
		} else {
			out += h.Name+"\n"
		}
	}
	for _, r := range d.rows {
		for i, data := range r {
			if i < len(r)-1 {
				out += tabifyName(data.Print())
			} else {
				out += data.Print()+"\n"
			}
		}
	}
	return out
}

func FilterCards(cards []*Card, predicate func(*Card) bool) []*Card {
	filteredCards := make([]*Card, 0)

	for _, c := range cards {
		if predicate(c) {
			filteredCards = append(filteredCards, c)
		}
	}
	return filteredCards
}


// func (m *TraitMap) Collect() CardCollection {
// 	a := make([]*Card, 0)
// 	for _, v := range *m {
// 		a = append(a, v...)
// 	}
// 	return CreateCollection(a)
// }

type DataCache struct {
	cardMap    	   *CardMap
	setMap     	   *SetMap
	typeMap    	   *TypeMap
	keywordMap 	   *KeywordMap
	traitMap   	   *TraitMap
	typeSynergyMap     *TypeMap
	traitSynergyMap    *TraitMap
	playAreaSynergyMap *PlayAreaSynergyMap
}

func (cache *DataCache) DumpStats() {
	for i, set := range *cache.setMap {
		fmt.Println("Set #"+strconv.Itoa(i)+": "+set[0].Name)
		// for _, c := range set {
		// 	fmt.Println("    "+c.Name)
		// }
	}

	for i, cards := range *cache.typeMap {
		fmt.Println("Type: "+CardTypeNames[i]+":"+strconv.Itoa(len(cards)))
	}

	traitCollection := CreateDataCollection("Trait", "Card Number", "Synergy Cards")
	for i, cards := range *cache.traitMap {
		traitCollection.AddRow(TraitNames[i], len(cards), len((*cache.traitSynergyMap)[i]))
	}
	traitCollection.Sort(func(i,j int) bool{return i > j}, 1, 2, 0)
	//traitCollection.FilterRow(func(r *DataRow) bool {return (*r)[2].IntValue() > 1})
	fmt.Print(traitCollection.Print())

	for i, cards := range *cache.keywordMap {
		fmt.Println("Keyword: "+KeywordNames[i]+":"+strconv.Itoa(len(cards)))
	}
}

func tabifyName(s string) string {
	if len(s) >= 8 {
		return s+"\t"
	}
	return s+"\t\t"
}

func AnalyzeDB(db []Card) []Card {
	cardMap        	   := make(CardMap)
	setMap         	   := make(SetMap)
	typeMap        	   := make(TypeMap)
	keywordMap     	   := make(KeywordMap)
	traitMap       	   := make(TraitMap)
	typeSynergyMap 	   := make(TypeMap)
	traitSynergyMap    := make(TraitMap)
	playAreaSynergyMap := make(PlayAreaSynergyMap, 0)

	for i, c := range db {
		// card definition uniqueness validation
		if cardMap[c.Number] != nil {
			panic("Card id "+strconv.Itoa(c.Number)+" is already present in DB, please merge them...")
			
		}
		cardPointer := &db[i]
		
		cardMap[c.Number] = cardPointer

		// set sanity validation
		for _, objSet := range c.ObjectiveSets {
			realIndex := objSet.CardSetNumber - 1
			if realIndex < 0 || realIndex > 5 {
				panic("Card "+strconv.Itoa(c.Number)+" has an invalid objective set card number: "+strconv.Itoa(objSet.CardSetNumber))
			} else if realIndex == 0 && c.Type.GetType() != CardType_Objective {
				panic("Trying to assing a non objective card as 1/6 for set #"+strconv.Itoa(objSet.SetId))
			}
			
			if setMap[objSet.SetId] == nil {
				setMap[objSet.SetId] = new(ObjectiveSetDB)
			} else if setMap[objSet.SetId][realIndex] != nil {
				panic("Cannot add card "+strconv.Itoa(c.Number)+" to set #"+strconv.Itoa(objSet.SetId)+" as card "+strconv.Itoa(objSet.CardSetNumber)+" / 6")
			}
			//fmt.Println("Adding "+c.Name+"in set "+strconv.Itoa(objSet.SetId))
			setMap[objSet.SetId][realIndex] = cardPointer
		}

		typeMap[c.Type.GetType()] = append(typeMap[c.Type.GetType()], &c)

		for _, ability := range c.Abilities {
			switch a := ability.(type) {
			case KeywordInterface: keywordMap[a.GetKeyword()] = append(keywordMap[a.GetKeyword()], cardPointer)
			case *CardTrait:       traitMap[a.Trait]          = append(traitMap[a.Trait], cardPointer)
			//case CardAbility: // todo map synergies
			}
		}


		for _, synergy := range c.GatherSynergies() {
			for i := 0 ; i < int(CardType_MAX) ; i++ {
				if synergy.IsSynergizingWithType(CardType(i)) {
					typeSynergyMap[CardType(i)] = append(typeSynergyMap[CardType(i)], cardPointer)
				}
			}

			for i := 0 ; i < int(Trait_MAX) ; i++ {
				if synergy.IsSynergizingWithTrait(CardTraitType(i)) {
					traitSynergyMap[CardTraitType(i)] = append(traitSynergyMap[CardTraitType(i)], cardPointer)
				}
			}

			if synergy.IsSynergizingWithPlayArea() {
				playAreaSynergyMap = append(playAreaSynergyMap, cardPointer)
			}
		}
	}

	cache := &DataCache{&cardMap, &setMap, &typeMap, &keywordMap, &traitMap, &typeSynergyMap, &traitSynergyMap, &playAreaSynergyMap}
	cache.DumpStats()
	
	return db
}

func CreateDB() []Card {
	return []Card{

		// SET 1 ------------------------------------------------------
		
		Card{ Name: "A Hero's Journey",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 2,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: nil,
			Health: 4,
			Quote: "\"You must learn the ways of the Force, if you are to come with me to Alderaan.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1,CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 2},

		Card{ Name: "Luke Skywalker",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 4,
			Ressources: 0,
			ForceIcons: 3,
			CardCombatIcons: CombatIcons(CombatIcon{2,0}, CombatIcon{0,0}, CombatIcon{1,1}),
			Abilities: AbilityList{
				Key(K_TargetedStrike),
				Trait(Trait_Character),
				Trait(Trait_ForceUser),
				Reaction("After your opponent's turn begins, remove 1 focus token from this unit.", nil)},
			Health: 3,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1,CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 92},

		Card{ Name: "Twi'lek Loyalist",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: CombatIcons(CombatIcon{0,1}, CombatIcon{0,1}, CombatIcon{0,0}),
			Abilities: AbilityList{Trait(Trait_Character)},
			Health: 1,
			Quote: "The Jedi once recruited from nearly every sentient species in the galaxy.",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 1, CardSetNumber: 3},
				ObjectiveSet{SetId: 4, CardSetNumber: 4},
			},
			Set: CardSet_Core,
			Number: 15},
		
		Card{ Name: "Jedi Lightsaber",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{TraitSynergy(Trait_ForceUser, true), TraitSynergy(Trait_ForceSensitive, true)}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
		        Abilities: AbilityList{
				Trait(Trait_Weapon),
				ConstantEffect("Enhanced Unit gains 1 Combat Damage and 1 Blast Damage.", nil)},
			Health: 0,
			Quote: "An elegant weapon for a more civilized age.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 102},

		Card{ Name: "Trust Your Feelings",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{TraitSynergy(Trait_Character, true)}),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
		        Abilities: AbilityList{
				Trait(Trait_Skill),
				Action("Focus this enhancement to remove 1 focus token from enhanced unit.", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 153},

		Card{ Name: "Dagobah Training Grounds",
			Faction: Faction_LightNeutral,
			Type: Enhancement(SynergyList{PlayAreaSynergy()}),
			Cost: 1,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Key(K_Limited),
				Trait(Trait_Dagobah),
				Trait(Trait_Location)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 1, CardSetNumber: 6},
				ObjectiveSet{SetId: 2, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 31},

		// SET 2 ------------------------------------------------------
		
		Card{ Name: "In You Must Go",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Dagobah),
				ConstantEffect("Reduce the cost of the first enhancement you play each turn by 1.",
					SynergyList{TypeSynergy(CardType_Enhancement, true)})},
			Health: 5,
			Quote: "\"What's in there?\"\n\"Only what you take with you.\"\n- The Empire Strikes Back",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 2, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 76},

		Card{ Name: "Yoda",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 3,
			Ressources: 0,
			ForceIcons: 5,
			CardCombatIcons: CombatIcons(CombatIcon{0,1}, CombatIcon{1,0}, CombatIcon{0,1}),
			Abilities: AbilityList{
				Key(K_Elite),
				Trait(Trait_Character),
				Trait(Trait_ForceUser),
				Reaction("After your opponent's turn begins, remove 1 focus token from this unit.", nil)},
			Health: 3,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 2, CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 166},

		Card{ Name: "Believer in the Old Ways",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: CombatIcons(CombatIcon{1,0}, CombatIcon{0,0}, CombatIcon{0,1}),
			Abilities: AbilityList{
				Trait(Trait_Character),
				Trait(Trait_ForceSensitive)},
			Health: 2,
			Quote: "\"The Jedi are extinct, their fire has gone out of the universe.\"\n-Grand Moff Wilhuff Tarkin, A New Hope",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 2, CardSetNumber: 3},
				ObjectiveSet{SetId: 7, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 154},
	
		Card{ Name: "Shii-Cho Training",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{TraitSynergy(Trait_ForceUser, true)}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Skill),
				Trait(Trait_LightSaberForm),
				ConstantEffect("Damage from enhanced unit's CombatDamage icon type may be divided among any number of participating enemy units.", nil)},
			Health: 0,
			Quote: "\"...let go your conscious self and act on instinct.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 2, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 122},

		// Dagobah Training Ground #31,

		Card{ Name: "Counter-Stroke",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Force),
				Trait(Trait_Control),
				Trait(Trait_Sense),
				Interrupt("When an event card is played, cancel its effect.", SynergyList{TypeSynergy(CardType_Event, false)})},
			Health: 0,
			Quote: "For those strong in the force, action and reaction are the same.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 2, CardSetNumber: 6}},
			Set: CardSet_Core,
			Number: 28},

		// SET 3 ------------------------------------------------------
		
		Card{ Name: "Forgotten Heroes",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Reaction("After you play a Force User unit, draw 1 card.",
					SynergyList{TraitSynergy(Trait_ForceUser, true)})},
			Health: 5,
			Quote: "\"The Force will be with you, always.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 95},

		Card{ Name: "Obi-Wan Kenobi",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 5,
			Ressources: 0,
			ForceIcons: 4,
			CardCombatIcons: CombatIcons(CombatIcon{1,0}, CombatIcon{1,1}, CombatIcon{0,1}),
			Abilities: AbilityList{
				Key(K_Elite),
				Trait(Trait_Character),
				Trait(Trait_ForceUser),
				ConstantEffect("While this unit is participating in an engagment, your opponent must place the first card of his edge stack faceup.", nil)},
			Health: 3,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 101},

		Card{ Name: "Jedi in Hiding",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: CombatIcons(CombatIcon{2,0}, CombatIcon{0,0}, CombatIcon{0,1}),
			Abilities: AbilityList{
				Trait(Trait_Character),
				Trait(Trait_ForceUser)},
			Health: 1,
			Quote: "\"For over a thousand generations the Jedi Knights were the guardians of peace and justice in the Old Republic. Before the dark times, before the Empire.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 3, CardSetNumber: 3},
			ObjectiveSet{SetId: 7, CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 84},
	
		Card{ Name: "Jedi Mind Trick",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Force),
				Trait(Trait_Control),
				Trait(Trait_Sense),
				Trait(Trait_Alter),
				Action("Place 1 focus token on a target Character or Creature unit. If the Balance of the Force is with the light side, place 2 focus tokens on that unit instead.",
					SynergyList{TraitSynergy(Trait_Character, false), TraitSynergy(Trait_Creature, false)})},
			Health: 0,
			Quote: "\"The Force can have a strong influence on the weak-minded.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 3, CardSetNumber: 4},
				ObjectiveSet{SetId: 7, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 85},

		Card{ Name: "Our Most Desperate Hour",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Event),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Action("Place 1 shield on a taret Character unit, even if that unit is already shielded.",
					SynergyList{TraitSynergy(Trait_Character, true)})},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 96},
		
		Card{ Name: "Heat of Battle",
			Faction: Faction_LightNeutral,
			Type: Fate(5),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Action("Deal 1 damage to a target participating enemy unit.", SynergyList{TypeSynergy(CardType_Unit, false)})},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 3, CardSetNumber: 6},
				ObjectiveSet{SetId: 18, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 65},

		// SET 4 ------------------------------------------------------
		
		Card{ Name: "A Journey to Dagobah",
			Faction: Faction_Jedi,
			Type: Objective(true),
			Cost: 0,
			Ressources: 2,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Dagobah),
				Interrupt("When this objective is destroyed, search your objective deck to choose your next objective and put it into play immediately. Shuffle your objective deck.", nil)},
			Health: 4,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 4, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 3},

		Card{ Name: "Red Five",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 3,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: CombatIcons(CombatIcon{0,0}, CombatIcon{0,0}, CombatIcon{3,0}),
			Abilities: AbilityList{
				Key(K_Elite),
				Trait(Trait_Vehicule),
				Trait(Trait_Fighter)},
			Health: 2,
				Quote: "\"Red Five, standing by.\"\nLuke Skywalker, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 4, CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 113},

		Card{ Name: "R2-D2",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Unit),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{Trait(Trait_Droid)},
			Health: 1,
				Quote: "Possessing a surprising amount of ingenuity for a droid, R2-D2 has rescued his friends from many tight situations.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 4, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 106},

			//Twi'lek Loyalist #15
	
		Card{ Name: "Double Strike",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: AbilityList{
					Reaction("After a Character unit is focused to strike, remove 1 focus token from that unit.",
						SynergyList{TraitSynergy(Trait_Character, true)})},
			Health: 0,
			Quote: "\"Not as clumsy or as random as a blaster.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 4, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 46},

		Card{ Name: "Target of Opportunity",
			Faction: Faction_LightNeutral,
			Type: Fate(9),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: AbilityList{
					Action("If you are the attacking player, deal 1 damage to the engaged objective.", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 4,  CardSetNumber: 6},
				ObjectiveSet{SetId: 18, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 133},

		// SET 5 ------------------------------------------------------
		
		Card{ Name: "The Secret of Yavin 4",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Yavin4),
				Interrupt("When 1 of your other objectives is engaged, your opponent engages this objective instead. [Limit once per turn.]",
					nil)},
			Health: 6,
			Quote: "The Rebels used the great temple on Yavin 4 to hide their command center, but never fully realized its mysterious history.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 5, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 144},

		Card{ Name: "C-3PO",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Unit),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Trait(Trait_Droid),
				Interrupt("When an event card is played, sacrifice this unit to cancel the effects of that event card.",
					SynergyList{TypeSynergy(CardType_Event, false)})},
			Health: 3,
			Quote: "\"Sir, if any of my circuits or gears will help, I'll gladly donate them.\"\n-C-3PO, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 5, CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 21},

		Card{ Name: "Guardian of Peace",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: CombatIcons(CombatIcon{0,1}, CombatIcon{0,0}, CombatIcon{0,0}),
			Abilities: AbilityList{
				Trait(Trait_Character),
				Trait(Trait_ForceSensitive),
				Key(K_Shielding),
				KeyProtect(Trait_Character),
			},
			Health: 2,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 5, CardSetNumber: 3},
				ObjectiveSet{SetId: 5, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 109},

		// Guardian of Peace #109
	
		Card{ Name: "Lightsaber Deflection",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Interrupt("When damage is dealt to a friendly non-Vehicle unit, deal 1 point of that damage to another target unit instead.",
					SynergyList{AccumulateSynergies(SynergyList{TypeSynergy(CardType_Unit, true), InvertSynergy(TraitSynergy(Trait_Vehicule, true))})})},
			Health: 0,
			Quote: "\"Good against remotes is on thing. Good against the living? That's something else.\"\n-han Solo, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 5, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 89},

		Card{ Name: "Twist of Fate",
			Faction: Faction_LightNeutral,
			Type: Fate(1),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Action("Cancel this edge battle and the card effects of all other fate cards just revealed. Discard both edge stacks and start a new edge battle..", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 5,  CardSetNumber: 6},
				ObjectiveSet{SetId: 18, CardSetNumber: 6}},
			Set: CardSet_Core,
			Number: 157},
		
		// SET 6 ------------------------------------------------------
		
		Card{ Name: "Last Minute Rescue",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
		Abilities: AbilityList{
				Trait(Trait_CloudCity),
				Reaction("After you refresh, remove 1 damage from a target unit.",
					SynergyList{TypeSynergy(CardType_Unit, true)})},
			Health: 5,
			Quote: "The Rebel alliance is outnumbered, outgunned, and commpletely overmatched. Yet still they have hope.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 118},

		Card{ Name: "Redemption",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 5,
			Ressources: 0,
			ForceIcons: 3,
			CardCombatIcons: CombatIcons(CombatIcon{2,0}, CombatIcon{0,0}, CombatIcon{0,2}),
		Abilities: AbilityList{
				Trait(Trait_Vehicule),
				Trait(Trait_CapitalShip),
				Interrupt("When a Character unit is destroyed, return it to its owner's hand instead of placing it in its owner's discard pile. [Limit once per turn.]",
					SynergyList{TraitSynergy(Trait_Character, true)})},
			Health: 4,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6,CardSetNumber: 2}},
			Set: CardSet_Core,
			Number: 115},

		Card{ Name: "Corellian Engineer",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Unit),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
		Abilities: AbilityList{
				Trait(Trait_Character),
				Trait(Trait_Engineer),
				Key(K_Shielding)},
			Health: 2,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 25},
		
		Card{ Name: "Return of the Jedi",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 3,
			Ressources: 0,
			ForceIcons: 3,
			CardCombatIcons: nil,
		Abilities: AbilityList{
				Action("Put a Force User unit into play from your discard pile.",
					SynergyList{TraitSynergy(Trait_ForceUser, true)})},
			Health: 0,
			Quote: "\"Luke, the Force runs strong in your family. Pass on what you have learned.\"\n-Yoda, Return of the Jedi",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 119},

		Card{ Name: "Emergency Repair",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Event),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
		        Abilities: AbilityList{Action("Remove all damage from a targe3t objective. [Play only during your turn]", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 50},

		Card{ Name: "Force Rejuvenation",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
	                Abilities: AbilityList{
				Trait(Trait_Force),
				Trait(Trait_Control),
				Action("Discard any number of tokens and enhancements from a target friendly Character unit.",
					SynergyList{TraitSynergy(Trait_Character, true)})},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 6, CardSetNumber: 6}},
			Set: CardSet_Core,
			Number: 61},

		// SET 7 ------------------------------------------------------
		
		Card{ Name: "Jedi Training",
			Faction: Faction_Jedi,
			Type: Objective(false),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
		Abilities: AbilityList{ConstantEffect("This objective contributes 1 Force icon to your side during the Force struggle.", nil)},
			Health: 5,
			Quote: "Jedi were once trained at the Jedi Temple on Coruscant. Now, the few that remain make due with what's available.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 7, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 86},

		// Jedi in Hiding #84
 	        // Believer in the Old Ways #154

		Card{ Name: "Ancient Monument",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{PlayAreaSynergy()}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 3,
			CardCombatIcons: nil,
		        Abilities: AbilityList{
				Trait(Trait_Location),
				ConstantEffect("This enhancement contributes 1 Force icon to your side during the Force struggle.", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 7,CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 6},

		// Jedi Mind Trick #85

		Card{ Name: "It Binds All Things",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Event),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: AbilityList{Action("Return the top card of your discard pile to your hand. If the Balance of the Force is with the light side, return the top 2 cards instead.", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 7, CardSetNumber: 6}},
			Set: CardSet_Core,
			Number: 80},

		// SET 18 ------------------------------------------------------
		
		Card{ Name: "Hit And Run",
			Faction: Faction_LightNeutral,
			Type: Objective(true),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: AbilityList{
				Reaction("After you win an edge battle as the attacker, deal 1 damage to the engaged objective. [Limit once per turn].", nil)},
			Health: 4,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 18, CardSetNumber: 1}},
			Set: CardSet_Core,
			Number: 68},

		Card{ Name: "Secret Informant",
			Faction: Faction_LightNeutral,
			Type: Type(CardType_Unit),
			Cost: 3,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: CombatIcons(CombatIcon{1,0}, CombatIcon{0,0}, CombatIcon{1,0}),
		        Abilities: AbilityList{
				Trait(Trait_Character),
				ConstantEffect("While this unit is participating in an engagement, you may resolve the effects of each fate card in your edge stack an additional time.",
					SynergyList{TypeSynergy(CardType_Fate, true)})},
			Health: 1,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{
				ObjectiveSet{SetId: 18,CardSetNumber: 2},
				ObjectiveSet{SetId: 18,CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 110},

		// Secret Informant #110
		// Heat of Battle #65
		// Target of Opportunity #133
		// Twist of Fate #157

	}
}
