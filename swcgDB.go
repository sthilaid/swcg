package swcg

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
			ObjectiveSets: []ObjectiveSet{ObjectiveSet{SetId: 1, CardSetNumber: 6}},
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
			Number: 92},

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
		Number: 28}}
}
