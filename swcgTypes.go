
package swcg

// Synergy Types --------------------------------------------------------------

type BaseSynergy struct {
	IsPositiveEff bool // if negative, is a synergy against the opponent's card
}
func (s BaseSynergy) IsPositiveEffect() bool {
	return s.IsPositiveEff
}
func (s BaseSynergy)IsSynergizingWithPlayArea()           bool { return false }
func (s BaseSynergy)IsSynergizingWith(*Card)          bool { return false }
func (s BaseSynergy)IsSynergizingWithType(CardType)       bool { return false }
func (s BaseSynergy)IsSynergizingWithTrait(CardTraitType) bool { return false}


// Type Synergy

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
func (syn *CardTypeSynergy)IsSynergizingWithType(t CardType) bool {
	return syn.Type == t
}

// Trait Synergy

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
func (syn *CardTraitSynergy)IsSynergizingWithTrait(t CardTraitType) bool {
	return syn.Trait == t
}

// Play Area Synergy

type PlayAreaSynergyType struct {
	BaseSynergy // not for the field, but for the methods
}
func PlayAreaSynergy() *PlayAreaSynergyType { return &PlayAreaSynergyType{} }
func (syn *PlayAreaSynergyType) IsSynergizingWith(*Card) bool {
	return false
}
func (syn *PlayAreaSynergyType) IsPositiveEffect() bool {
	return true
}
func (syn *PlayAreaSynergyType)IsSynergizingWithPlayArea() bool {
	return true
}

// Meta Synergies

// Inversion 
type InvertedSynergyType struct {
	synergy SynergyInterface
}
func InvertSynergy(s SynergyInterface) *InvertedSynergyType {
	return &InvertedSynergyType{synergy: s}
}
func (syn *InvertedSynergyType) IsSynergizingWith(c *Card) bool {
	return !syn.synergy.IsSynergizingWith(c)
}
func (syn *InvertedSynergyType) IsSynergizingWithType(t CardType) bool {
	return !syn.synergy.IsSynergizingWithType(t)
}
func (syn *InvertedSynergyType) IsSynergizingWithTrait(t CardTraitType) bool {
	return !syn.synergy.IsSynergizingWithTrait(t)
}
func (syn *InvertedSynergyType)IsSynergizingWithPlayArea() bool {
	return false
}
func (syn *InvertedSynergyType) IsPositiveEffect() bool {
	return syn.synergy.IsPositiveEffect()
}

// Accumulation
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

func (syn *AccumulationSynergyType) IsSynergizingWithType(t CardType) bool {
	for _, s := range syn.synergies {
		if s.IsSynergizingWithType(t) {
			return true
		}
	}
	return false
}
func (syn *AccumulationSynergyType) IsSynergizingWithTrait(t CardTraitType) bool {
	for _, s := range syn.synergies {
		if s.IsSynergizingWithTrait(t) {
			return true
		}
	}
	return false
}

func (syn *AccumulationSynergyType)IsSynergizingWithPlayArea() bool {
	return false
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

// Options
type OptionalSynergyType struct {
	AccumulationSynergyType
}
func SynergyOptions(ss SynergyList) *OptionalSynergyType {
	return &OptionalSynergyType{AccumulationSynergyType: *AccumulateSynergies(ss)}
}
func (syn *OptionalSynergyType) IsSynergizingWith(c *Card) bool {
	for _, s := range syn.synergies {
		if s.IsSynergizingWith(c) {
			return true
		}
	}
	return false
}

// Synergy Interface

type SynergyInterface interface {
	IsPositiveEffect() bool
	IsSynergizingWithPlayArea() bool
	IsSynergizingWith(*Card) bool
	IsSynergizingWithType(CardType) bool
	IsSynergizingWithTrait(CardTraitType) bool
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
var CardTypeNames [CardType_MAX]string = [CardType_MAX]string {
	"Unit",
	"Event",
	"Objective",
	"Fate",
	"Enhancement",
}

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
	Faction_LightNeutral   CardFaction = iota
	Faction_Sith           CardFaction = iota
	Faction_ImperialNavy   CardFaction = iota
	Faction_ScumAndVillany CardFaction = iota
	Faction_DarkNeutral    CardFaction = iota
	Faction_MAX            CardFaction = iota
)
var FactionNames [Faction_MAX]string = [Faction_MAX]string {
	"Jedi",
	"RebelAliance",
	"Smugglers",
	"LightNeutral",
	"Sith",
	"ImperialNavy",
	"ScumAndVillany",
	"DarkNeutral",
}

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

var AbilityNames[AbilityType_MAX]string = [AbilityType_MAX]string {
	"Action",
	"Reaction",
	"Interrupt",
	"ConstantEffect",
	"Keyword",
	"Trait",
}

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
var KeywordNames [K_MAX]string = [K_MAX]string {
	"Edge",
	"Elite",
	"Limited",
	"NoEnhancement",
	"Protect",
	"Shielding",
	"TargetedStrike",
}

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
	ProtectedTrait CardTraitType
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

var TraitNames [Trait_MAX]string = [Trait_MAX]string {
	"Character",
	"Vehicule",
	"Force",
	"ForceUser",
	"ForceSensitive",
	"Weapon",
	"Skill",
	"Dagobah",
	"Location",
	"LightSaberForm",
	"Control",
	"Sense",
	"Alter",
	"Creature",
	"Fighter",
	"Droid",
	"Yavin4",
	"CloudCity",
	"CapitalShip",
	"Engineer",
}
	
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

var SetNames [CardSet_MAX]string = [CardSet_MAX]string {
	"Core",
}


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

func (c *Card) GatherSynergies() SynergyList {
	list := make(SynergyList, 0)

	if enhancementType, ok := c.Type.(*EnhancementCardType) ; ok {
		list = append(list, enhancementType.Synergies...)
	}

	for _, ability := range c.Abilities {
		switch castedAbility := ability.(type) {
		case *ProtectKeywordType:
			list = append(list, TraitSynergy(castedAbility.ProtectedTrait, true))
		case *CardAbility:
			list = append(list, castedAbility.Synergies...)
		}
	}
	return list
}
