package enums

type PlayerSide byte

const (
	Usec PlayerSide = 1 << iota
	Bear
	Savage
)
