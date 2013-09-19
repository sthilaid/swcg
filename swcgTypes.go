
package swcg

// Synergy Types --------------------------------------------------------------

type CardTypeSynergy struct {
	Type CardTypeInterface
}
func TypeSynergy(t CardTypeInterface) *CardTypeSynergy {
	return &CardTypeSynergy{Type: t}
}
func (syn *CardTypeSynergy) IsSynergizingWith(card *Card) bool {
	return syn.Type.GetType() == card.Type.GetType()
}

type CardTraitSynergy struct {
	Trait CardTrait
}
func TraitSynergy(trait CardTrait) *CardTraitSynergy {
	return &CardTraitSynergy{Trait: trait}
}
func (syn *CardTraitSynergy) IsSynergizingWith(card *Card) bool {
	for _, cardTrait := range card.Traits {
		if syn.Trait == cardTrait {
			return true
		}
	}
	return false
}

type PlayAreaSynergyType struct {}
func PlayAreaSynergy() *PlayAreaSynergyType { return &PlayAreaSynergyType{} }
func (syn *PlayAreaSynergyType) IsSynergizingWith(*Card) bool {
	return false
}

type SynergyInterface interface {
	IsSynergizingWith(card *Card) bool
}
type SynergyList []SynergyInterface


// Card Types -----------------------------------------------------------------

type CardType int
const (
	CardType_Unit  	     CardType = iota
	CardType_Event 	     CardType = iota
	CardType_Objective   CardType = iota
	CardType_Fate        CardType = iota
	CardType_Enhancement CardType = iota
	CardType_MAX         CardType = iota
)

type SimpleCardType struct {
	Type CardType
}
func Type(t CardType) *SimpleCardType { return &SimpleCardType{Type: t} }
func (t *SimpleCardType) GetType() CardType {
	return t.Type
}

type EnhancementCardType struct {
	SimpleCardType
	Synergies []SynergyInterface
}
func Enhancement(Synergies SynergyList) *EnhancementCardType {
	return &EnhancementCardType{SimpleCardType: SimpleCardType{Type: CardType_Enhancement},
	                            Synergies: Synergies}
}

type FateCardType struct {
	SimpleCardType
	EdgeBattlePriority int
}
func Fate(priority int) *FateCardType {
	return &FateCardType{SimpleCardType: SimpleCardType{Type: CardType_Fate}, EdgeBattlePriority: priority}
}

type CardTypeInterface interface {
	GetType() CardType
}

// Keywords  ------------------------------------------------------------------

type CardKeywordType int
const (
	K_Edge    	   CardKeywordType = iota
	K_Elte    	   CardKeywordType = iota
	K_Limited 	   CardKeywordType = iota
	K_NoEnhancement    CardKeywordType = iota
	K_Protect   	   CardKeywordType = iota
	K_Shielding 	   CardKeywordType = iota
	K_TargetedStrike   CardKeywordType = iota
	K_MAX              CardKeywordType = iota
)
type SimpleKeyword struct {
	K CardKeywordType
}
func Key(K CardKeywordType) *SimpleKeyword { return &SimpleKeyword{K: K} }
func (k *SimpleKeyword) GetKeyword() CardKeywordType { return k.K }

type ComplexKeyword struct {
	SimpleKeyword
	V int
}
func KeyEdge(n int) *ComplexKeyword { return &ComplexKeyword{SimpleKeyword: SimpleKeyword{K: K_Edge}, V: n} }

type KeywordInterface interface {
	GetKeyword() CardKeywordType
}

// Traits  --------------------------------------------------------------------

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
	//...
	Trait_MAX             CardTrait = iota
)

// Factions  ------------------------------------------------------------------

type CardFaction int
const (
	Faction_Jedi           CardFaction = iota
	Faction_RebelAliance   CardFaction = iota
	Faction_Smugglers      CardFaction = iota
	Faction_Sith           CardFaction = iota
	Faction_ImperialNavy   CardFaction = iota
	Faction_ScumAndVillany CardFaction = iota
	Faction_Neutral        CardFaction = iota
	Faction_MAX            CardFaction = iota
)

// Combat Icons  --------------------------------------------------------------

type CombatIcon [2]int
type CombatIcons struct {
	CombatDamage CombatIcon
	Tactics      CombatIcon
	BlastDamage  CombatIcon
}

// Actions   --------------------------------------------------------------

type ActionType int
const (
	ActionType_Action      ActionType = iota
	ActionType_Reaction    ActionType = iota
	ActionType_Enhancement ActionType = iota
	//...
	ActionType_MAX         ActionType = iota
)
type CardAction struct {
	Type           ActionType
	Description    string
	Synergies      SynergyList
}

// Sets   ---------------------------------------------------------------------

type CardSet int
const (
	CardSet_Core CardSet = iota
	CardSet_MAX  CardSet = iota
)

// Cards   --------------------------------------------------------------------

type Card struct {
	Name            string
	Faction         CardFaction
	Type            CardTypeInterface
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
