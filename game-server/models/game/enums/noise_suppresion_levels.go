package enums

type NoiseSuppressionLevels int

const (
	NoiseSuppressionDisabled NoiseSuppressionLevels = -1
	NoiseSuppressionLow      NoiseSuppressionLevels = iota
	NoiseSuppressionModerate
	NoiseSuppressionHigh
	NoiseSuppressionVeryHigh
)
