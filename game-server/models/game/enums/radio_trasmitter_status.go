package enums

type RadioTransmitterStatus int

const (
	NotInitialized RadioTransmitterStatus = iota
	NoRadioTransmitter
	Red
	Green
	Yellow
)
