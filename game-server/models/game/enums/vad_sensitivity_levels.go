package enums

type VadSensitivityLevels int

const (
	VadLowSensitivity VadSensitivityLevels = iota
	VadMediumSensitivity
	VadHighSensitivity
	VadVeryHighSensitivity
)
