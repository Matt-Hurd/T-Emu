package enums

type DamageType int

const (
	Undefined DamageType = 1 << iota
	Fall
	Explosion
	Barbed
	Flame
	GrenadeFragment
	Impact
	Existence
	Medicine
	Bullet
	Melee
	Landmine
	Sniper
	Blunt
	LightBleeding
	HeavyBleeding
	Dehydration
	Exhaustion
	RadExposure
	Stimulator
	Poison
	LethalToxin
	Btr
)
