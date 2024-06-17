package core

type VadSensitivityLevels int

const (
	VadLowSensitivity VadSensitivityLevels = iota
	VadMediumSensitivity
	VadHighSensitivity
	VadVeryHighSensitivity
)

type NoiseSuppressionLevels int

const (
	NoiseSuppressionDisabled NoiseSuppressionLevels = -1
	NoiseSuppressionLow      NoiseSuppressionLevels = iota
	NoiseSuppressionModerate
	NoiseSuppressionHigh
	NoiseSuppressionVeryHigh
)

type AudioQuality int

const (
	AudioQualityLow AudioQuality = iota
	AudioQualityMedium
	AudioQualityHigh
)

type FrameSize int

const (
	FrameSizeTiny  FrameSize = -1
	FrameSizeSmall FrameSize = iota
	FrameSizeMedium
	FrameSizeLarge
)

type EMemberCategory int

const (
	MemberCategoryDefault EMemberCategory = iota
	MemberCategoryDeveloper
	MemberCategoryUniqueId
	MemberCategoryTrader                        = 4
	MemberCategoryGroup                         = 8
	MemberCategorySystem                        = 16
	MemberCategoryChatModerator                 = 32
	MemberCategoryChatModeratorWithPermanentBan = 64
	MemberCategoryUnitTest                      = 128
	MemberCategorySherpa                        = 256
	MemberCategoryEmissary                      = 512
)

type ENetLogsLevel int

const (
	NetLogsLevelNone ENetLogsLevel = iota
	NetLogsLevelMinimum
	NetLogsLevelNormal
	NetLogsLevelMaximum
)
