package enums

type HandsControllerType byte

const (
	HandsControllerTypeNone HandsControllerType = iota
	Empty
	Firearm
	Meds
	Grenade
	Knife
	QuickGrenade
	QuickKnife
	UsableItem
	QuickUseItem
)
