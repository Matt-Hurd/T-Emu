package math

import (
	"github.com/g3n/engine/math32"
)

func ScaleShortToFloat(value int16, minTarget, maxTarget float32) float32 {
	return ScaleShortToFloatClamped(value, -32768, 32767, minTarget, maxTarget)
}

func ScaleShortToFloatClamped(value, minValue, maxValue int16, minTarget, maxTarget float32) float32 {
	num := maxTarget - minTarget
	num2 := uint16(maxValue - minValue)
	num3 := uint16(math32.ClampInt(int(value), int(minValue), int(maxValue)) - int(minValue))
	return minTarget + (float32(num3)/float32(num2))*num
}

func ScaleByteToFloat(value byte, minTarget, maxTarget float32) float32 {
	return ScaleByteToFloatClamped(value, 0, 255, minTarget, maxTarget)
}

func ScaleByteToFloatClamped(value, minValue, maxValue byte, minTarget, maxTarget float32) float32 {
	num := maxTarget - minTarget
	num2 := float32(maxValue - minValue)
	num3 := float32(math32.ClampInt(int(value), int(minValue), int(maxValue)) - int(minValue))
	return minTarget + num3/num2*num
}

func ScaleFromVector2Byte(value []byte, minTarget, maxTarget float32) Vector2 {
	return ScaleFromVector2ByteClamped(value, 0, 255, minTarget, maxTarget)
}

func ScaleFromVector2ByteClamped(value []byte, minValue, maxValue byte, minTarget, maxTarget float32) Vector2 {
	return Vector2{
		math32.Vector2{
			X: ScaleByteToFloatClamped(value[0], minValue, maxValue, minTarget, maxTarget),
			Y: ScaleByteToFloatClamped(value[1], minValue, maxValue, minTarget, maxTarget),
		},
	}
}

func ScaleFromVector3Short(value []int16, minTarget, maxTarget float32) Vector3 {
	return ScaleFromVector3ShortClamped(value, -32768, 32767, minTarget, maxTarget)
}

func ScaleFromVector3ShortClamped(value []int16, minValue, maxValue int16, minTarget, maxTarget float32) Vector3 {
	return Vector3{
		math32.Vector3{
			X: ScaleShortToFloatClamped(value[0], minValue, maxValue, minTarget, maxTarget),
			Y: ScaleShortToFloatClamped(value[1], minValue, maxValue, minTarget, maxTarget),
			Z: ScaleShortToFloatClamped(value[2], minValue, maxValue, minTarget, maxTarget),
		},
	}
}

func ScaleByteToInt(value byte, minTarget, maxTarget int32) int32 {
	return ScaleByteToIntClamped(value, 0, 255, minTarget, maxTarget)
}

func ScaleByteToIntClamped(value, minValue, maxValue byte, minTarget, maxTarget int32) int32 {
	num := maxTarget - minTarget
	num2 := int32(maxValue - minValue)
	num3 := int32(math32.ClampInt(int(value), int(minValue), int(maxValue)) - int(minValue))
	return minTarget + num3/num2*num
}

func ScaleFloatToShort(value, minValue, maxValue float32) int16 {
	return ScaleFloatToShortClamped(value, minValue, maxValue, -32768, 32767)
}

func ScaleFloatToShortClamped(value, minValue, maxValue float32, minTarget, maxTarget int16) int16 {
	num := maxTarget - minTarget
	num2 := float32(maxTarget - minTarget)
	num3 := math32.Clamp(value, minValue, maxValue) - minValue
	return minTarget + int16(num3/num2*float32(num))
}

func ScaleFloatToByte(value, minValue, maxValue float32) byte {
	return ScaleFloatToByteClamped(value, minValue, maxValue, 0, 255)
}

func ScaleFloatToByteClamped(value, minValue, maxValue float32, minTarget, maxTarget byte) byte {
	num := maxTarget - minTarget
	num2 := float32(maxTarget - minTarget)
	num3 := math32.Clamp(value, minValue, maxValue) - minValue
	return minTarget + byte(num3/num2*float32(num))
}

func ScaleToVector2Byte(value Vector2, minTarget, maxTarget float32) []byte {
	return ScaleToVector2ByteClamped(value, minTarget, maxTarget, 0, 255)
}

func ScaleToVector2ByteClamped(value Vector2, minTarget, maxTarget float32, minValue, maxValue byte) []byte {
	return []byte{
		ScaleFloatToByteClamped(value.X, minTarget, maxTarget, minValue, maxValue),
		ScaleFloatToByteClamped(value.Y, minTarget, maxTarget, minValue, maxValue),
	}
}

func ScaleToVector3Short(value Vector3, minTarget, maxTarget float32) []int16 {
	return ScaleToVector3ShortClamped(value, minTarget, maxTarget, -32768, 32767)
}

func ScaleToVector3ShortClamped(value Vector3, minTarget, maxTarget float32, minValue, maxValue int16) []int16 {
	return []int16{
		ScaleFloatToShortClamped(value.X, minTarget, maxTarget, minValue, maxValue),
		ScaleFloatToShortClamped(value.Y, minTarget, maxTarget, minValue, maxValue),
		ScaleFloatToShortClamped(value.Z, minTarget, maxTarget, minValue, maxValue),
	}
}

func ScaleIntToByte(value, minValue, maxValue int32) byte {
	return ScaleIntToByteClamped(value, minValue, maxValue, 0, 255)
}

func ScaleIntToByteClamped(value, minValue, maxValue int32, minTarget, maxTarget byte) byte {
	num := maxTarget - minTarget
	num2 := float32(maxValue - minValue)
	num3 := float32(math32.ClampInt(int(value), int(minValue), int(maxValue))) - float32(minValue)
	return minTarget + byte(num3/num2*float32(num))
}
