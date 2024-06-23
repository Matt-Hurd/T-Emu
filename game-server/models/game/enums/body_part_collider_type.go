package enums

type BodyPartColliderType int16

const (
	BodyPartColliderNone BodyPartColliderType = iota - 1
	HeadCommon
	RibcageUp
	Pelvis BodyPartColliderType = iota
	LeftUpperArm
	LeftForearm
	RightUpperArm
	RightForearm
	LeftThigh
	LeftCalf
	RightThigh
	RightCalf
	ParietalHead
	BackHead
	Ears
	Eyes
	Jaw
	NeckFront
	NeckBack
	RightSideChestUp
	LeftSideChestUp
	SpineTop
	SpineDown
	PelvisBack
	RightSideChestDown
	LeftSideChestDown
	RibcageLow
)
