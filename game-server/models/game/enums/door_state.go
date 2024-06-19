package enums

type EDoorState byte

const (
	None EDoorState = iota
	DoorStateLocked
	Shut
	Open        = 4
	Interacting = 8
	Breaching   = 16
)
