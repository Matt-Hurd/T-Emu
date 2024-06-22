package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type LimitedWriter struct {
	Buffer    *bytes.Buffer
	bitBuffer uint64
	bitCount  int
	byteCount int
	bitPos    int
}

func NewLimitedWriter(Buffer *bytes.Buffer) *LimitedWriter {
	return &LimitedWriter{
		Buffer:    Buffer,
		bitBuffer: 0,
		bitCount:  0,
		byteCount: 0,
		bitPos:    0,
	}
}

func (lw *LimitedWriter) BitsWritten() int {
	return lw.bitCount
}

func (lw *LimitedWriter) BytesWritten() int {
	return lw.byteCount
}

func (lw *LimitedWriter) WriteBits(value uint32, bits int) error {
	value &= (1 << bits) - 1
	lw.bitBuffer |= uint64(value) << (64 - lw.bitPos - bits)
	lw.bitPos += bits
	lw.bitCount += bits

	for lw.bitPos >= 32 {
		toWrite := uint32(lw.bitBuffer >> 32)
		err := binary.Write(lw.Buffer, binary.LittleEndian, toWrite)
		if err != nil {
			return err
		}
		lw.bitBuffer <<= 32
		lw.bitPos -= 32
		lw.byteCount += 4
	}

	return nil
}

func (lw *LimitedWriter) WriteBytes(bytes []byte) error {
	for _, b := range bytes {
		err := lw.WriteBits(uint32(b), 8)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lw *LimitedWriter) WriteByteAlign() {
	if lw.bitPos%8 != 0 {
		lw.WriteBits(0, 8-lw.bitPos%8)
	}
}

func (lw *LimitedWriter) WriteLimitedInt32(value, min, max int) error {
	bits := bitsRequired(min, max)
	return lw.WriteBits(uint32(value-min), int(bits))
}

func (lw *LimitedWriter) WriteLimitedFloat(value, min, max, res float64) error {
	delta := (max - min) / res
	bits := bitsRequired(0, int(delta))
	quantizedValue := uint32((value - min) / res)
	return lw.WriteBits(quantizedValue, int(bits))
}

func (lw *LimitedWriter) WriteLimitedString(value string, min, max rune) error {
	if err := lw.Write(false); err != nil {
		return err
	}
	lw.WriteByteAlign()
	length := int32(len(value))
	if err := lw.Write(length); err != nil {
		return err
	}
	bits := bitsRequired(int(min), int(max))

	for _, r := range value {
		if err := lw.WriteBits(uint32(r-min), int(bits)); err != nil {
			return err
		}
	}

	return nil
}

// func (lw *LimitedWriter) WriteString(value string, min, max rune) error {
// 	if err := lw.Write(false); err != nil {
// 		return err
// 	}
// 	lw.WriteByteAlign()
// 	length := int32(len(value))
// 	if err := lw.Write(length); err != nil {
// 		return err
// 	}
// 	for _, r := range value {
// 		if err := lw.WriteBits(uint32(r-min), 8); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (lw *LimitedWriter) WriteEnum(min, max, value int) error {
	bits := int(bitsRequired(min, max))
	return lw.WriteBits(uint32(value-min), bits)
}

func (lw *LimitedWriter) Write(value interface{}) error {
	switch v := value.(type) {
	case bool:
		if v {
			return lw.WriteBits(1, 1)
		}
		return lw.WriteBits(0, 1)
	case int:
		return lw.WriteBits(uint32(v), 32)
	case uint:
		return lw.WriteBits(uint32(v), 32)
	case byte:
		return lw.WriteBits(uint32(v), 8)
	case int16:
		return lw.WriteBits(uint32(v), 16)
	case uint16:
		return lw.WriteBits(uint32(v), 16)
	case int32:
		return lw.WriteBits(uint32(v), 32)
	case uint32:
		return lw.WriteBits(v, 32)
	case int64:
		if err := lw.WriteBits(uint32(v), 32); err != nil {
			return err
		}
		return lw.WriteBits(uint32(v>>32), 32)
	case uint64:
		if err := lw.WriteBits(uint32(v), 32); err != nil {
			return err
		}
		return lw.WriteBits(uint32(v>>32), 32)
	case float32:
		return lw.WriteBits(math.Float32bits(v), 32)
	case float64:
		bits := math.Float64bits(v)
		if err := lw.WriteBits(uint32(bits), 32); err != nil {
			return err
		}
		return lw.WriteBits(uint32(bits>>32), 32)
	// case string:
	// 	length := uint16(len(v))
	// 	if err := lw.Write(length); err != nil {
	// 		return err
	// 	}
	// 	return lw.WriteString(v, 0, 255)
	default:
		return fmt.Errorf("unsupported type %T", v)
	}
}

func (lw *LimitedWriter) FlushBits() error {
	if lw.bitPos > 0 {
		toWrite := uint32(lw.bitBuffer >> 32)
		err := binary.Write(lw.Buffer, binary.LittleEndian, toWrite)
		if err != nil {
			return err
		}
		lw.bitBuffer <<= 32
		lw.bitPos -= 32
		lw.byteCount += 4
	}
	return nil
}
