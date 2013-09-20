
package swcg

// Synergy Types --------------------------------------------------------------

type BaseSynergy struct {
	IsPositiveEff bool // if negative, is a synergy against the opponent's card
}
func (s BaseSynergy) IsPositiveEffect() bool {
	return s.IsPositiveEff
}

type CardTypeSynergy struct {
	BaseSynergy
	Type CardType
}
func TypeSynergy(t CardType, isPositive bool) *CardTypeSynergy {
	return &CardTypeSynergy{BaseSynergy: BaseSynergy{isPositive}, Type: t}
}
func (syn *CardTypeSynergy) IsSynergizingWith(card *Card) bool {
	return syn.Type == card.Type.GetType()
}

type CardTraitSynergy struct {
	BaseSynergy
	Trait CardTraitType
}
func TraitSynergy(trait CardTraitType, isPositive bool) *CardTraitSynergy {
	return &CardTraitSynergy{BaseSynergy: BaseSynergy{isPositive}, Trait: trait}
}
func (syn *CardTraitSynergy) IsSynergizingWith(card *Card) bool {
	for _, ability := range card.Abilities {
		cardTrait, ok := ability.(*CardTrait)
		if ok && syn.Trait == cardTrait.Trait {
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
func (syn *PlayAreaSynergyType) IsPositiveEffect() bool {
	return true
}

type InvertedSynergyType struct {
	synergy SynergyInterface
}
func InvertSynergy(s SynergyInterface) *InvertedSynergyType {
	return &InvertedSynergyType{synergy: s}
}
func (syn *InvertedSynergyType) IsSynergizingWith(c *Card) bool {
	return !syn.synergy.IsSynergizingWith(c)
}
func (syn *InvertedSynergyType) IsPositiveEffect() bool {
	return syn.synergy.IsPositiveEffect()
}

type AccumulationSynergyType struct {
	synergies SynergyList
}
func AccumulateSynergies(ss SynergyList) *AccumulationSynergyType {
	return &AccumulationSynergyType{synergies: ss}
}
func (syn *AccumulationSynergyType) IsSynergizingWith(c *Card) bool {
	for _, s := range syn.synergies {
		if !s.IsSynergizingWith(c) {
			return false
		}
	}
	return true
}
func (syn *AccumulationSynergyType) IsPositiveEffect() bool {
	var isPositive bool
	for i, s := range syn.synergies {
		if i == 0 {
			isPositive = s.IsPositiveEffect()
		} else if isPositive != s.IsPositiveEffect() {
			panic("found inconsistent positive effect declaration in accumulation synergy...")
		}
	}
	return isPositive
}

type OptionalSynergyType struct {
	synergies SynergyList
}
func SynergyOptions(ss SynergyList) *OptionalSynergyType {
	return &OptionalSynergyType{synergies: ss}
}
func (syn *OptionalSynergyType) IsSynergizingWith(c *Card) bool {
	for _, s := range syn.synergies {
		if s.IsSynergizingWith(c) {
			return true
		}
	}
	return false
}
func (syn *OptionalSynergyType) IsPositiveEffect() bool {
	var isPositive bool
	for i, s := range syn.synergies {
		if i == 0 {
			isPositive = s.IsPositiveEffect()
		} else if isPositive != s.IsPositiveEffect() {
			panic("found inconsistent positive effect declaration in optinal synergy...")
		}
	}
	return isPositive
}



type SynergyInterface interface {
	IsSynergizingWith(card *Card) bool
	IsPositiveEffect() bool
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
func Type(t CardType) *SimpleCardType {
	if t == CardType_Enhancement || t == CardType_Fate || t == CardType_Objective {
		panic("Enhancement, Fate and Objective card types shouldn't be constructed with the Type function...")
	}
	return &SimpleCardType{Type: t}
}
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

type ObjectiveCardType struct {
	SimpleCardType
	OnlyAvailableToFaction bool
}
func Objective(OnlyAvailableToFaction bool) *ObjectiveCardType {
	return &ObjectiveCardType{SimpleCardType: SimpleCardType{Type: CardType_Objective}, OnlyAvailableToFaction: OnlyAvailableToFaction}
}

type CardTypeInterface interface {
	GetType() CardType
}

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
type CardCombatIcons struct {
	CombatDamage CombatIcon
	Tactics      CombatIcon
	BlastDamage  CombatIcon
}
func CombatIcons(combat CombatIcon, tactics CombatIcon, blast CombatIcon) *CardCombatIcons {
	return &CardCombatIcons{CombatDamage: combat, Tactics: tactics, BlastDamage: blast}
}

// Abilities   ----------------------------------------------------------------

type AbilityType int
const (
	AbilityType_Action          AbilityType = iota
	AbilityType_Reaction        AbilityType = iota
	AbilityType_Interrupt       AbilityType = iota
	AbilityType_ConstantEffect  AbilityType = iota
	AbilityType_Keyword         AbilityType = iota
	AbilityType_Trait           AbilityType = iota
	AbilityType_MAX             AbilityType = iota
)

type BaseAbility struct {
	Type           AbilityType
}
func (a *BaseAbility) GetType() AbilityType {
	return a.Type
}

type CardAbility struct {
	BaseAbility
	Description string
	Synergies   SynergyList
}
func Ability(t AbilityType, desc string, syn SynergyList) *CardAbility {
	if t == AbilityType_Keyword || t == AbilityType_Trait {
		panic("Keywords and Traits are not basic abilities...")
	}
	return &CardAbility{BaseAbility: BaseAbility{Type: t}, Description: desc, Synergies: syn}
}
func Action(desc string, syn SynergyList) *CardAbility {
	return Ability(AbilityType_Action, desc, syn)
}
func Reaction(desc string, syn SynergyList) *CardAbility {
	return Ability(AbilityType_Reaction, desc, syn)
}
func Interrupt(desc string, syn SynergyList) *CardAbility {
	return Ability(AbilityType_Interrupt, desc, syn)
}
func ConstantEffect(desc string, syn SynergyList) *CardAbility {
	return Ability(AbilityType_ConstantEffect, desc, syn)
}

type AbilityInterface interface {
	GetType() AbilityType
}
type AbilityList []AbilityInterface


// Keywords  ------------------------------------------------------------------

type CardKeywordType int
const (
	K_Edge    	   CardKeywordType = iota
	K_Elite    	   CardKeywordType = iota
	K_Limited 	   CardKeywordType = iota
	K_NoEnhancement    CardKeywordType = iota
	K_Protect   	   CardKeywordType = iota
	K_Shielding 	   CardKeywordType = iota
	K_TargetedStrike   CardKeywordType = iota
	K_MAX              CardKeywordType = iota
)
type SimpleKeyword struct {
	BaseAbility
	K CardKeywordType
}
func Key(K CardKeywordType) *SimpleKeyword {
	return &SimpleKeyword{BaseAbility: BaseAbility{AbilityType_Keyword}, K: K}
}
func (k *SimpleKeyword) GetKeyword() CardKeywordType { return k.K }

type ComplexKeyword struct {
	SimpleKeyword
	V int
}
func KeyEdge(n int) *ComplexKeyword { return &ComplexKeyword{SimpleKeyword: *Key(K_Edge), V: n} }

type ProtectKeywordType struct {
	SimpleKeyword
	ProtectedTrait CardTraitType  // !TODO! count in synergy value
}
func KeyProtect(protectedTrait CardTraitType) *ProtectKeywordType {
	return &ProtectKeywordType{SimpleKeyword: *Key(K_Protect), ProtectedTrait: protectedTrait}
}

type KeywordInterface interface {
	GetKeyword() CardKeywordType
}

// Traits  --------------------------------------------------------------------

type CardTraitType int
const (
	Trait_Character       CardTraitType = iota
	Trait_Vehicule        CardTraitType = iota
	Trait_Force           CardTraitType = iota
	Trait_ForceUser       CardTraitType = iota
	Trait_ForceSensitive  CardTraitType = iota
	Trait_Weapon          CardTraitType = iota
	Trait_Skill           CardTraitType = iota
	Trait_Dagobah         CardTraitType = iota
	Trait_Location        CardTraitType = iota
	Trait_LightSaberForm  CardTraitType = iota
	Trait_Control         CardTraitType = iota
	Trait_Sense           CardTraitType = iota
	Trait_Alter           CardTraitType = iota
	Trait_Creature        CardTraitType = iota
	Trait_Fighter         CardTraitType = iota
	Trait_Droid           CardTraitType = iota
	Trait_Yavin4          CardTraitType = iota
	Trait_CloudCity       CardTraitType = iota
	Trait_CapitalShip     CardTraitType = iota
	Trait_Engineer        CardTraitType = iota
	//...
	Trait_MAX             CardTraitType = iota
)

type CardTrait struct {
	BaseAbility
	Trait CardTraitType
}
func Trait(t CardTraitType) *CardTrait {
	return &CardTrait{BaseAbility: BaseAbility{Type: AbilityType_Trait}, Trait: t}
}

// Sets   ---------------------------------------------------------------------

type ObjectiveSet struct {
	SetId int
	CardSetNumber int // from 1 to 6, 1 is always the objective card
}

type CardSetType int
const (
	CardSet_Core CardSetType = iota
	CardSet_MAX  CardSetType = iota
)

// Cards   --------------------------------------------------------------------

type Card struct {
	Name            string
	Faction         CardFaction
	Type            CardTypeInterface
	Cost            int
	Ressources      int
	ForceIcons      int
	CardCombatIcons *CardCombatIcons
	Abilities       []AbilityInterface
	Health          int
	Quote           string
	ObjectiveSets   []ObjectiveSet
	Set             CardSetType
	Number          int
}
