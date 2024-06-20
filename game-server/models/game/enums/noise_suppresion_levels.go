package enums

type NoiseSuppressionLevels int

const (
	NoiseSuppressionDisabled NoiseSuppressionLevels = iota - 1
	NoiseSuppressionLow
	NoiseSuppressionModerate
	NoiseSuppressionHigh
	NoiseSuppressionVeryHigh
)
