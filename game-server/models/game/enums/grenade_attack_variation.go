package enums

type GrenadeAttackVariation byte

const (
	GrenadeAttackVariationNone GrenadeAttackVariation = iota
	High
	Low
	QuickHigh
	QuickLow
)
