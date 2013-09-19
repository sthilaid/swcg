package swcg

import "strconv"

type ObjectiveSetDB [6]*Card

func AnalyzeDB(db []Card) []Card{
	cardMap := make(map[int]*Card)
	setMap  := make(map[int]*ObjectiveSetDB)
	
	for _, c := range db {
		if cardMap[c.Number] != nil {
			panic("Card id "+strconv.Itoa(c.Number)+" is already present in DB, please merge them...")
			
		}
		cardMap[c.Number] = &c

		for _, objSet := range c.ObjectiveSets {
			realIndex := objSet.CardSetNumber - 1
			if realIndex < 0 || realIndex > 5 {
				panic("Card "+strconv.Itoa(c.Number)+" has an invalid objective set card number: "+strconv.Itoa(objSet.CardSetNumber))
			}
			
			if setMap[objSet.SetId] == nil {
				setMap[objSet.SetId] = new(ObjectiveSetDB)
			} else if setMap[objSet.SetId][realIndex] != nil {
				panic("Cannot add card "+strconv.Itoa(c.Number)+" to set #"+strconv.Itoa(objSet.SetId)+" as card "+strconv.Itoa(objSet.CardSetNumber)+" / 6")
			} 
			setMap[objSet.SetId][realIndex] = &c
		}
	}
	return db
}

func CreateDB() []Card {
	return []Card{

		// SET 1 ------------------------------------------------------
		
		Card{ Name: "A Hero's Journey",
			Faction: Faction_Jedi,
			Type: Type(CardType_Objective),
			Cost: 0,
			Ressources: 2,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: nil,
			Health: 0,
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
			Abilities: []AbilityInterface{
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
			Abilities: []AbilityInterface{Trait(Trait_Character)},
			Health: 1,
			Quote: "The Jedi once recruited from nearly every sentient species in the galaxy.",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 15},
		
		Card{ Name: "Jedi Lightsaber",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{TraitSynergy(Trait_ForceUser, true), TraitSynergy(Trait_ForceSensitive, true)}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
		        Abilities: []AbilityInterface{
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
		        Abilities: []AbilityInterface{
				Trait(Trait_Skill),
				Action("Focus this enhancement to remove 1 focus token from enhanced unit.", nil)},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 153},

		Card{ Name: "Dagobah Training Grounds",
			Faction: Faction_Neutral,
			Type: Enhancement(SynergyList{PlayAreaSynergy()}),
			Cost: 1,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
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
			Type: Type(CardType_Objective),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
				Trait(Trait_Dagobah),
				ConstantEffect("Reduce the cost of the first enhancement you play each turn by 1.",
					SynergyList{TypeSynergy(CardType_Enhancement, true)})},
			Health: 0,
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
			Abilities: []AbilityInterface{
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
			Abilities: []AbilityInterface{
				Trait(Trait_Character),
				Trait(Trait_ForceSensitive)},
			Health: 2,
			Quote: "\"The Jedi are extinct, their fire has gone out of the universe.\"\n-Grand Moff Wilhuff Tarkin, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 2, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 154},
	
		Card{ Name: "Shii-Cho Training",
			Faction: Faction_Jedi,
			Type: Enhancement(SynergyList{TraitSynergy(Trait_ForceUser, true)}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
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
			Abilities: []AbilityInterface{
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
			Type: Type(CardType_Objective),
			Cost: 0,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
				Reaction("After you play a Force User unit, draw 1 card.",
					SynergyList{TraitSynergy(Trait_ForceUser, true)})},
			Health: 0,
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
			Abilities: []AbilityInterface{
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
			Abilities: []AbilityInterface{
				Trait(Trait_Character),
				Trait(Trait_ForceUser)},
			Health: 1,
			Quote: "\"For over a thousand generations the Jedi Knights were the guardians of peace and justice in the Old Republic. Before the dark times, before the Empire.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 3}},
			Set: CardSet_Core,
			Number: 84},
	
		Card{ Name: "Jedi Mind Trick",
			Faction: Faction_Jedi,
			Type: Type(CardType_Event),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
				Trait(Trait_Force),
				Trait(Trait_Control),
				Trait(Trait_Sense),
				Trait(Trait_Alter),
				Action("Place 1 focus token on a target Character or Creature unit. If the Balance of the Force is with the light side, place 2 focus tokens on that unit instead.",
					SynergyList{TraitSynergy(Trait_Character, false), TraitSynergy(Trait_Creature, false)})},
			Health: 0,
			Quote: "\"The Force can have a strong influence on the weak-minded.\"\n-Obi-Wan Kenobi, A New Hope",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 4}},
			Set: CardSet_Core,
			Number: 85},

		Card{ Name: "Our Most Desperate Hour",
			Faction: Faction_Neutral,
			Type: Type(CardType_Event),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
				Action("Place 1 shield on a taret Character unit, even if that unit is already shielded.",
					SynergyList{TraitSynergy(Trait_Character, true)})},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 5}},
			Set: CardSet_Core,
			Number: 96},
		
		Card{ Name: "Heat of Battle",
			Faction: Faction_Neutral,
			Type: Fate(5),
			Cost: 0,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Abilities: []AbilityInterface{
				Action("Deal 1 damage to a target participating enemy unit.", SynergyList{TypeSynergy(CardType_Unit, false)})},
			Health: 0,
			Quote: "",
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 3, CardSetNumber: 6}},
			Set: CardSet_Core,
		Number: 65}}
}
