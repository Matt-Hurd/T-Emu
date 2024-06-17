package enums

type ENetLogsLevel int

const (
	NetLogsLevelNone ENetLogsLevel = iota
	NetLogsLevelMinimum
	NetLogsLevelNormal
	NetLogsLevelMaximum
)
