package enums

type ELampState byte

const (
	TurningOn ELampState = iota
	TurningOff
	On
	Off
	Destroyed
	ConstantFlickering
	SmoothOff
)
