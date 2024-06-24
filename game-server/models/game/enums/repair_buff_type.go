package enums

type RepairBuffType int32

const (
	WeaponSpread RepairBuffType = iota
	DamageReduction
	MalfunctionProtections
	WeaponDamage
	ArmorEfficiency
	DurabilityImprovement
)
