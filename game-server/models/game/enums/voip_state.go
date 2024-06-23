package enums

type VoipState byte

const (
	NotAvailable VoipState = iota
	Available
	VoipStateOff
	Banned
	MicrophoneFail
)
