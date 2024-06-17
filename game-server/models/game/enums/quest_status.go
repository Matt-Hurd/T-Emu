package enums

type EQuestStatus int

const (
	Locked EQuestStatus = iota
	AvailableForStart
	Started
	AvailableForFinish
	Success
	Fail
	FailRestartable
	MarkedAsFailed
	Expired
	AvailableAfter
)
