package enums

type TagStatus uint32

const (
	TagStatusUnaware TagStatus = 1 << iota
	TagStatusAware
	TagStatusCombat
	TagStatusSolo
	TagStatusCoop
	TagStatusBear
	TagStatusUsec
	TagStatusScav
	TagStatusTargetSolo
	TagStatusTargetMultiple
	TagStatusHealthy
	TagStatusInjured
	TagStatusBadlyInjured
	TagStatusDying
	TagStatusBirdeye
	TagStatusKnight
	TagStatusBigPipe
)
