package enums

type ArmorPlateCollider int16

const (
	Plate_Granit_SAPI_chest ArmorPlateCollider = 1 << iota
	Plate_Granit_SAPI_back
	Plate_Granit_SSAPI_side_left_high
	Plate_Granit_SSAPI_side_left_low
	Plate_Granit_SSAPI_side_right_high
	Plate_Granit_SSAPI_side_right_low
	Plate_Korund_chest
	Plate_6B13_back
	Plate_Korund_side_left_high
	Plate_Korund_side_left_low
	Plate_Korund_side_right_high
	Plate_Korund_side_right_low
)
