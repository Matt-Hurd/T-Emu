package enums

type RadioTransmitterStatus byte

const (
	NotInitialized RadioTransmitterStatus = iota
	NoRadioTransmitter
	Red
	Green
	Yellow
)
