package enums

type MalfunctionState byte

const (
	MalfunctionStateNone MalfunctionState = iota
	Misfire
	Jam
	HardSlide
	SoftSlide
	Feed
)
