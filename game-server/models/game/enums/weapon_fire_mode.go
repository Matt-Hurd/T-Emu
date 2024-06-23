package enums

type FireMode byte

const (
	fullauto FireMode = iota
	single
	doublet
	burst
	doubleaction
	semiauto
)
