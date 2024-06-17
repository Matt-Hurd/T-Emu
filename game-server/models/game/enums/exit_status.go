package enums

type ExitStatus int

const (
	Survived ExitStatus = iota
	Killed
	Left
	Runner
	MissingInAction
)
