
package swcg

type CardType int
const (
	CardType_Unit  	     CardType = iota
	CardType_Event 	     CardType = iota
	CardType_Objective   CardType = iota
	CardType_Fate        CardType = iota
	CardType_Enhancement CardType = iota
)

type SimpleCardType struct {
	Type CardType
}
func Type(t CardType) *SimpleCardType { return &SimpleCardType{Type: t} }
func (t *SimpleCardType) GetType() CardType {
	return t.Type
}

type EnhancementCardType struct {
	Type CardType
	Synergies []CardTrait
}
func Enhancement(Synergies []CardTrait) *EnhancementCardType { return &EnhancementCardType{Type: CardType_Enhancement, Synergies: Synergies} }
func (t *EnhancementCardType) GetType() CardType {
	return t.Type
}

type CardTypeInt interface {
	GetType() CardType
}

type CardKeywordType int
const (
	K_Edge    	   CardKeywordType = iota
	K_Elte    	   CardKeywordType = iota
	K_Limited 	   CardKeywordType = iota
	K_NoEnhancement  CardKeywordType = iota
	K_Protect   	   CardKeywordType = iota
	K_Shielding 	   CardKeywordType = iota
	K_TargetedStrike CardKeywordType = iota
)
type SimpleKeyword struct {
	K CardKeywordType
}
func Key(K CardKeywordType) *SimpleKeyword { return &SimpleKeyword{K: K} }
func (k *SimpleKeyword) GetKeyword() CardKeywordType { return k.K }

type ComplexKeyword struct {
	K CardKeywordType
	V int
}
func (k *ComplexKeyword) GetKeyword() CardKeywordType { return k.K }
func KeyEdge(n int) *ComplexKeyword { return &ComplexKeyword{K: K_Edge, V: n} }

type KeywordInterface interface {
	GetKeyword() CardKeywordType
}


type CardTrait int
const (
	Trait_Character       CardTrait = iota
	Trait_Vehicule        CardTrait = iota
	Trait_ForceUser       CardTrait = iota
	Trait_ForceSensitive  CardTrait = iota
	Trait_Weapon          CardTrait = iota
	Trait_Skill           CardTrait = iota
	Trait_Dagobah         CardTrait = iota
	Trait_Location        CardTrait = iota
	Trait_PlayArea        CardTrait = iota // dummy trait, more a game concept
	//...
)

type CardFaction int
const (
	Faction_Jedi           CardFaction = iota
	Faction_RebelAliance   CardFaction = iota
	Faction_Smugglers      CardFaction = iota
	Faction_Sith           CardFaction = iota
	Faction_ImperialNavy   CardFaction = iota
	Faction_ScumAndVillany CardFaction = iota
	Faction_Neutral        CardFaction = iota
)

type CombatIcon [2]int
type CombatIcons struct {
	CombatDamage CombatIcon
	Tactics      CombatIcon
	BlastDamage  CombatIcon
}

type ActionType int
const (
	ActionType_Action      ActionType = iota
	ActionType_Reaction    ActionType = iota
	ActionType_Enhancement ActionType = iota
	//...
)
type CardAction struct {
	Type           ActionType
	Description    string
	SynergicTraits []CardTrait
}

type CardSet int
const (
	CardSet_Core CardSet = iota
)

type Card struct {
	Name            string
	Faction         CardFaction
	Type            CardTypeInt
	Cost            int
	Ressources      int
	ForceIcons      int
	CardCombatIcons *CombatIcons
	Keywords        []KeywordInterface
	Traits          []CardTrait
	Actions         []CardAction
	Health          int
	Quote           string
	Block           int
	BlockId         int
	Set             CardSet
	Number          int
}
