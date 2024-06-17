package enums

type FrameSize int

const (
	FrameSizeTiny  FrameSize = -1
	FrameSizeSmall FrameSize = iota
	FrameSizeMedium
	FrameSizeLarge
)
