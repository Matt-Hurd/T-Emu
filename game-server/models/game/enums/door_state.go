package enums

type EDoorState byte

const (
	None EDoorState = iota
	DoorStateLocked
	Shut
	Open        EDoorState = 4
	Interacting EDoorState = 8
	Breaching   EDoorState = 16
)
