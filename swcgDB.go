package swcg

func CreateDB() []Card {
	return []Card{ 
		Card{ Name: "A Hero's Journey",
			Faction: Faction_Jedi,
			Type: Type(CardType_Objective),
			Cost: 0,
			Ressources: 2,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Keywords: []KeywordInterface{},
			Traits: []CardTrait{},
			Actions: []CardAction{},
			Health: 0,
			Quote: "\"You must learn the ways of the Force, if you are to come with me to Alderaan.\"\n-Obi-Wan Kenobi, A New Hope",
			Block: 1,
			BlockId: 1,
			Set: CardSet_Core,
			Number: 2},

		Card{ Name: "Luke Skywalker",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 4,
			Ressources: 0,
			ForceIcons: 3,
			CardCombatIcons: &CombatIcons{CombatDamage: CombatIcon{2,0}, Tactics: CombatIcon{0,0}, BlastDamage: CombatIcon{1,1}},
			Keywords: []KeywordInterface{Key(K_TargetedStrike)},
			Traits: []CardTrait{Trait_Character, Trait_ForceUser},
			Actions: []CardAction{CardAction{Type: ActionType_Reaction, Description: "After your opponent's turn begins, remove 1 focus token from this unit.", SynergicTraits: nil}},
			Health: 3,
			Quote: "",
			Block: 1,
			BlockId: 2,
			Set: CardSet_Core,
			Number: 92},

		Card{ Name: "Twi'lek Loyalist",
			Faction: Faction_Jedi,
			Type: Type(CardType_Unit),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: &CombatIcons{CombatDamage: CombatIcon{0,1}, Tactics: CombatIcon{0,1}, BlastDamage: CombatIcon{0,0}},
			Keywords: []KeywordInterface{},
			Traits: []CardTrait{Trait_Character},
			Actions: []CardAction{},
			Health: 1,
			Quote: "The Jedi once recruited from nearly every sentient species in the galaxy.",
			Block: 1,
			BlockId: 3,
			Set: CardSet_Core,
			Number: 15},
	
			Card{ Name: "Jedi Lightsaber",
			Faction: Faction_Jedi,
			Type: Enhancement([]CardTrait{Trait_ForceUser, Trait_ForceSensitive}),
			Cost: 1,
			Ressources: 0,
			ForceIcons: 2,
			CardCombatIcons: nil,
			Keywords: []KeywordInterface{},
			Traits: []CardTrait{Trait_Weapon},
			Actions: []CardAction{CardAction{Type: ActionType_Enhancement, Description: "Enhanced Unit gains 1 Combat Damage and 1 Blast Damage.", SynergicTraits: nil}},
			Health: 0,
			Quote: "An elegant weapon for a more civilized age.",
			Block: 1,
			BlockId: 4,
			Set: CardSet_Core,
			Number: 102},

			Card{ Name: "Trust Your Feelings",
			Faction: Faction_Jedi,
			Type: Enhancement([]CardTrait{Trait_Character}),
			Cost: 2,
			Ressources: 0,
			ForceIcons: 1,
			CardCombatIcons: nil,
			Keywords: []KeywordInterface{},
			Traits: []CardTrait{Trait_Skill},
			Actions: []CardAction{CardAction{Type: ActionType_Action, Description: "Focus this enhancement to remove 1 focus token from enhanced unit.", SynergicTraits: nil}},
			Health: 0,
			Quote: "",
			Block: 1,
			BlockId: 5,
			Set: CardSet_Core,
			Number: 153},

			Card{ Name: "Dagobah Training Grounds",
			Faction: Faction_Neutral,
			Type: Enhancement([]CardTrait{Trait_PlayArea}),
			Cost: 1,
			Ressources: 1,
			ForceIcons: 0,
			CardCombatIcons: nil,
			Keywords: []KeywordInterface{Key(K_Limited)},
			Traits: []CardTrait{Trait_Dagobah, Trait_Location},
			Actions: []CardAction{},
			Health: 0,
			Quote: "",
			Block: 1,
			BlockId: 6,
			Set: CardSet_Core,
		Number: 31}}

}
