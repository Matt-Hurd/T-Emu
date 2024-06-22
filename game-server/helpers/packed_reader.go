package helpers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type LimitedReader struct {
	ByteBuffer *bytes.Buffer
	byteSlice  []byte
	int0       int
	int1       int
	int2       int
	int3       int
	int4       int
	bool0      bool
	firstRead  bool
	ulong0     uint64
}

func NewLimitedReader(buffer *bytes.Buffer) *LimitedReader {
	length := buffer.Len()
	// Add 4 bytes to the buffer to prevent overflow
	buffer.Write(make([]byte, 4))
	return &LimitedReader{
		ByteBuffer: buffer,
		byteSlice:  buffer.Bytes(),
		int0:       length / 4,
		int1:       length / 4 * 32,
		int2:       0,
		int3:       0,
		int4:       0,
		firstRead:  true,
		bool0:      false,
	}
}

func (lr *LimitedReader) IsOverflow() bool {
	return lr.bool0
}

func (lr *LimitedReader) BitsRead() int {
	return lr.int2
}

func (lr *LimitedReader) BitsCount() int {
	return lr.int1
}

func (lr *LimitedReader) BytesRead() int {
	return (lr.int4 + func() int {
		if lr.int3 > 0 {
			return 1
		}
		return 0
	}()) * 4
}

func (lr *LimitedReader) TotalBytes() int {
	return lr.int1 * 8
}

func (lr *LimitedReader) ReadBits(bits int) (uint32, error) {
	if lr.firstRead {
		var tmp uint32
		if err := binary.Read(lr.ByteBuffer, binary.LittleEndian, &tmp); err != nil {
			return 0, err
		}
		lr.ulong0 = uint64(tmp)
		lr.firstRead = false
	}

	if lr.int2+bits > lr.int1 {
		lr.bool0 = true
		return 0, errors.New("overflow")
	}

	lr.int2 += bits
	if lr.int3+bits < 32 {
		lr.ulong0 <<= bits
		lr.int3 += bits
	} else {
		lr.int4++
		num := 32 - lr.int3
		num2 := bits - num
		lr.ulong0 <<= num
		var temp uint32
		if err := binary.Read(lr.ByteBuffer, binary.LittleEndian, &temp); err != nil {
			return 0, err
		}
		lr.ulong0 |= uint64(temp)
		lr.ulong0 <<= num2
		lr.int3 = num2
	}

	result := uint32(lr.ulong0 >> 32)
	lr.ulong0 &= 0xFFFFFFFF
	return result, nil
}

func (lr *LimitedReader) ReadBytes(destination []byte, destinationStartIndex int, bytesCount int) error {
	if lr.firstRead {
		if err := binary.Read(lr.ByteBuffer, binary.LittleEndian, &lr.ulong0); err != nil {
			return err
		}
		lr.firstRead = false
	}

	if lr.int2+bytesCount*8 > lr.int1 {
		copy(destination[destinationStartIndex:], make([]byte, bytesCount))
		lr.bool0 = true
		return errors.New("overflow")
	}

	num := (4 - lr.int3/8) % 4
	if num > bytesCount {
		num = bytesCount
	}

	for i := 0; i < num; i++ {
		bit, err := lr.ReadBits(8)
		if err != nil {
			return err
		}
		destination[destinationStartIndex+i] = byte(bit)
	}

	if num == bytesCount {
		return nil
	}

	num2 := (bytesCount - num) / 4
	if num2 > 0 {
		length := num2 * 4
		sourceIndex := lr.int4 * 4
		copy(destination[destinationStartIndex+num:], lr.byteSlice[sourceIndex:sourceIndex+length])
		lr.int2 += num2 * 32
		lr.int4 += num2
		var temp uint32
		if err := binary.Read(lr.ByteBuffer, binary.LittleEndian, &temp); err != nil {
			return err
		}
		lr.ulong0 = uint64(temp)
	}

	num3 := num + num2*4
	num4 := bytesCount - num3
	for j := 0; j < num4; j++ {
		bit, err := lr.ReadBits(8)
		if err != nil {
			return err
		}
		destination[destinationStartIndex+num3+j] = byte(bit)
	}

	return nil
}

func (lr *LimitedReader) ReadAlign() {
	num := lr.int2 % 8
	if num != 0 {
		lr.ReadBits(8 - num)
	}
}

func (br *LimitedReader) ReadLimitedInt32(min, max int) (int, error) {
	bits := bitsRequired(min, max)
	value, err := br.ReadBits(int(bits))
	if err != nil {
		return 0, err
	}
	return int(value) + min, nil
}

func (br *LimitedReader) ReadLimitedFloat(min, max, res float64) (float64, error) {
	delta := (max - min) / res
	bits := bitsRequired(0, int(delta))
	quantizedValue, err := br.ReadBits(int(bits))
	if err != nil {
		return 0, err
	}
	return min + float64(quantizedValue)*res, nil
}

func (lr *LimitedReader) ReadLimitedString(min, max rune) (string, error) {

	var null bool
	err := lr.Read(&null)
	if err != nil {
		return "", err
	}
	if null {
		return "", nil
	}
	lr.ReadAlign()
	var length int32
	err = lr.Read(&length)
	if err != nil {
		return "", err
	}
	array := make([]rune, length)
	bits := bitsRequired(int(min), int(max))

	for i := 0; i < int(length); i++ {
		num5, err := lr.ReadBits(int(bits))
		if err != nil {
			return "", err
		}
		array[i] = rune(num5) + min
	}

	return string(array), nil
}

func (br *LimitedReader) ReadString(min, max rune) (string, error) {
	var null bool
	err := br.Read(&null)
	if err != nil {
		return "", err
	}
	if null {
		return "", nil
	}
	br.ReadAlign()
	var length int32
	err = br.Read(&length)
	if err != nil {
		return "", err
	}
	fmt.Printf("length: %d\n", length)
	runes := make([]rune, length)
	for i := 0; i < int(length); i++ {
		charValue, err := br.ReadBits(8)
		if err != nil {
			return "", err
		}
		runes[i] = rune(charValue) + min
	}
	return string(runes), nil
}

func (br *LimitedReader) ReadEnum(min, max int) (int, error) {
	bits := int(bitsRequired(min, max))
	value, err := br.ReadBits(bits)
	if err != nil {
		return 0, err
	}
	return int(value) + min, nil
}

func (br *LimitedReader) Read(value interface{}) error {
	switch v := value.(type) {
	case *bool:
		val, err := br.ReadBits(1)
		if err != nil {
			return err
		}
		*v = val == 1
	case *int:
		val, err := br.ReadBits(32)
		if err != nil {
			return err
		}
		*v = int(val)
	case *uint:
		val, err := br.ReadBits(32)
		if err != nil {
			return err
		}
		*v = uint(val)
	case *byte:
		val, err := br.ReadBits(8)
		if err != nil {
			return err
		}
		*v = byte(val)
	case *int16:
		val, err := br.ReadBits(16)
		if err != nil {
			return err
		}
		*v = int16(val)
	case *uint16:
		val, err := br.ReadBits(16)
		if err != nil {
			return err
		}
		*v = uint16(val)
	case *int32:
		val, err := br.ReadBits(32)
		if err != nil {
			return err
		}
		*v = int32(val)
	case *uint32:
		val, err := br.ReadBits(32)
		if err != nil {
			return err
		}
		*v = val
	case *int64:
		var lower, upper uint32
		var err error
		if lower, err = br.ReadBits(32); err != nil {
			return err
		}
		if upper, err = br.ReadBits(32); err != nil {
			return err
		}
		*v = int64(lower) | (int64(upper) << 32)
	case *uint64:
		var lower, upper uint32
		var err error
		if lower, err = br.ReadBits(32); err != nil {
			return err
		}
		if upper, err = br.ReadBits(32); err != nil {
			return err
		}
		*v = uint64(lower) | (uint64(upper) << 32)
	case *float32:
		val, err := br.ReadBits(32)
		if err != nil {
			return err
		}
		*v = math.Float32frombits(val)
	case *float64:
		var lower, upper uint32
		var err error
		if lower, err = br.ReadBits(32); err != nil {
			return err
		}
		if upper, err = br.ReadBits(32); err != nil {
			return err
		}
		*v = math.Float64frombits(uint64(lower) | (uint64(upper) << 32))
	case *string:
		var length uint16
		if err := br.Read(&length); err != nil {
			return err
		}
		str, err := br.ReadString(0, 255)
		if err != nil {
			return err
		}
		*v = str
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

func bitsRequired(min, max int) int {
	if min == max {
		return 0
	}
	return int(Log2(uint(max-min)) + 1)
}
func Log2(value uint) uint {
	num := value | value>>1
	num2 := num | num>>2
	num3 := num2 | num2>>4
	num4 := num3 | num3>>8
	result := (num4 | num4>>16) >> 1
	return smethod_0(result)
}

func smethod_0(x uint) uint {
	num := x - (x >> 1 & 1431655765)
	num2 := (num >> 2 & 858993459) + (num & 858993459)
	num3 := (num2 >> 4) + num2&252645135
	num4 := num3 + (num3 >> 8)
	return num4 + (num4>>16)&63
}
