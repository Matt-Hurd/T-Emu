package enums

type PhysicalCondition int

const (
	PhysicalConditionNone PhysicalCondition = 0
	OnPainkillers         PhysicalCondition = 1 << (iota - 1)
	LeftLegDamaged
	RightLegDamaged
	ProneDisabled
	LeftArmDamaged
	RightArmDamaged
	Tremor
	UsingMeds
	HealingLegs
	JumpDisabled
	SprintDisabled
	ProneMovementDisabled
	Panic
)
