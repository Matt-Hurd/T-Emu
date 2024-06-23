package helpers

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"
)

func ReadString(buffer *bytes.Buffer, result *string) error {
	var byteCount uint16
	err := binary.Read(buffer, binary.LittleEndian, &byteCount)
	if err != nil {
		return err
	}

	if byteCount == 0 {
		return nil
	}

	byteData := make([]byte, byteCount)
	n, err := buffer.Read(byteData)
	if err != nil {
		return err
	}

	if n != int(byteCount) {
		return errors.New("read: incorrect number of bytes read")
	}

	*result = string(byteData)

	return nil
}

func read7BitEncodedInt(buffer *bytes.Buffer) (int, error) {
	num := 0
	num2 := 0
	for {
		if num2 == 35 {
			return 0, errors.New("Format_Bad7BitInt32")
		}
		b, err := buffer.ReadByte()
		if err != nil {
			return 0, err
		}
		num |= int(b&0x7F) << num2
		num2 += 7
		if (b & 0x80) == 0 {
			break
		}
	}
	return num, nil
}

func ReadUTF16String(buffer *bytes.Buffer, result *string) error {
	byteCount, err := read7BitEncodedInt(buffer)
	if err != nil {
		return err
	}
	byteData := make([]byte, byteCount)
	n, err := buffer.Read(byteData)
	if err != nil {
		return err
	}

	if n != int(byteCount) {
		return errors.New("read: incorrect number of bytes read")
	}

	*result = string(byteData)

	// if err != nil {
	// 	return err
	// }
	// if num2 < 0 {
	// 	return fmt.Errorf("invalid string length: %d", num2)
	// }
	// if num2 == 0 {
	// 	return nil
	// }

	// u16s := make([]uint16, num2)
	// for i := 0; i < num2; i++ {
	// 	var u16 uint16
	// 	err := binary.Read(buffer, binary.LittleEndian, &u16)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u16s[i] = u16
	// }

	// runes := utf16.Decode(u16s)
	// var stringBuilder bytes.Buffer
	// for _, r := range runes {
	// 	buf := make([]byte, utf8.UTFMax)
	// 	size := utf8.EncodeRune(buf, r)
	// 	stringBuilder.Write(buf[:size])
	// }
	// *result = stringBuilder.String()
	return nil
}

func ReadBool(buffer *bytes.Buffer, result *bool) error {
	b, err := buffer.ReadByte()
	if err != nil {
		return err
	}
	*result = b == 1
	return nil
}

func ReadBytesAndSize(buffer *bytes.Buffer, result *[]byte) error {
	var length uint16
	err := binary.Read(buffer, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	*result = make([]byte, length)
	_, err = buffer.Read(*result)
	if err != nil {
		return err
	}
	return nil
}

func ReadUInt32(buffer *bytes.Buffer, result *uint32) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadInt16(buffer *bytes.Buffer, result *int16) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadInt32(buffer *bytes.Buffer, result *int32) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadInt64(buffer *bytes.Buffer, result *int64) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadUInt64(buffer *bytes.Buffer, result *uint64) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadFloat64(buffer *bytes.Buffer, result *float64) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadFloat32(buffer *bytes.Buffer, result *float32) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadUInt16(buffer *bytes.Buffer, result *uint16) error {
	return binary.Read(buffer, binary.LittleEndian, result)
}

func ReadByte(buffer *bytes.Buffer, result *byte) error {
	b, err := buffer.ReadByte()
	if err != nil {
		return err
	}
	*result = b
	return nil
}

func Int64ToTime(int64Time int64) time.Time {
	const ticksToNanoseconds = 100
	const unixEpochTicks = 621355968000000000
	ticks := int64Time & 0x3FFFFFFFFFFFFFFF

	nanoseconds := (ticks - unixEpochTicks) * ticksToNanoseconds
	gameDateTime := time.Unix(0, nanoseconds).UTC()
	return gameDateTime
}

func ReadDateTime(buffer *bytes.Buffer, result *time.Time) error {
	var binaryDate int64
	err := binary.Read(buffer, binary.LittleEndian, &binaryDate)
	if err != nil {
		return err
	}

	*result = Int64ToTime(binaryDate)
	return nil
}

func DecompressZlib(compressed []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = io.Copy(&decompressed, reader)
	if err != nil {
		return nil, err
	}

	return decompressed.Bytes(), nil
}

func ReadPackedUInt64(buf *bytes.Buffer, ret *uint64) error {
	b, err := buf.ReadByte()
	if err != nil {
		return err
	}

	var result uint64
	var bytesToRead int

	switch {
	case b < 241:
		*ret = uint64(b)
		return nil
	case b >= 241 && b <= 248:
		bytesToRead = 1
		result = 240 + 256*(uint64(b)-241)
	case b == 249:
		bytesToRead = 2
		result = 2288
	case b == 250:
		bytesToRead = 3
	case b == 251:
		bytesToRead = 4
	case b == 252:
		bytesToRead = 5
	case b == 253:
		bytesToRead = 6
	case b == 254:
		bytesToRead = 7
	case b == 255:
		bytesToRead = 8
	default:
		return errors.New("invalid first byte")
	}

	for i := 0; i < bytesToRead; i++ {
		nextByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		result += uint64(nextByte) << (8 * uint(i))
	}

	*ret = result
	return nil
}

func ReadPackedUInt32(buf *bytes.Buffer, out *uint32) error {
	firstByte, err := buf.ReadByte()
	if err != nil {
		*out = 0
		return err
	}

	if firstByte <= 240 {
		*out = uint32(firstByte)
		return nil
	}
	if firstByte <= 248 {
		secondByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		*out = 240 + uint32(firstByte-241)*256 + uint32(secondByte)
		return nil
	}
	if firstByte == 249 {
		secondByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		thirdByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		*out = 2288 + uint32(secondByte)*256 + uint32(thirdByte)
		return nil
	}
	if firstByte == 250 {
		secondByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		thirdByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		fourthByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		*out = uint32(secondByte) + (uint32(thirdByte) << 8) + (uint32(fourthByte) << 16)
		return nil
	}
	if firstByte == 251 {
		secondByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		thirdByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		fourthByte, err := buf.ReadByte()
		if err != nil {
			return err
		}
		fifthByte, err := buf.ReadByte()
		if err != nil {
			return err
		}

		*out = uint32(secondByte) + (uint32(thirdByte) << 8) + (uint32(fourthByte) << 16) + (uint32(fifthByte) << 24)
		return nil
	}

	return fmt.Errorf("invalid first byte: %d", firstByte)
}

func ReadMongoId(buffer *bytes.Buffer, id *string) error {
	var timestamp uint32
	var counter uint64
	if err := ReadUInt32(buffer, &timestamp); err != nil {
		return err
	}
	if err := ReadUInt64(buffer, &counter); err != nil {
		return err
	}
	*id = fmt.Sprintf("%08x%06x", timestamp, counter)
	return nil
}
