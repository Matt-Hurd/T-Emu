package helpers

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"os"
	"time"
	"unicode/utf16"
)

const (
	ticksToNanoseconds = 100
	unixEpochTicks     = 621355968000000000
)

func WriteString(buffer *bytes.Buffer, value string) error {
	if value == "" {
		err := binary.Write(buffer, binary.LittleEndian, uint16(0))
		return err
	}

	byteCount := len(value)
	if byteCount >= 32768 {
		return errors.New("Serialize(string) too long")
	}

	err := binary.Write(buffer, binary.LittleEndian, uint16(byteCount))
	if err != nil {
		return err
	}

	_, err = buffer.Write([]byte(value))
	return err
}

func UInt32ToBytes(value uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)
	return buf.Bytes()
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func WriteBool(buffer *bytes.Buffer, value bool) error {
	var b byte
	if value {
		b = 1
	} else {
		b = 0
	}
	return buffer.WriteByte(b)
}

func WriteBytesAndSize(buffer *bytes.Buffer, value []byte) error {
	length := uint16(len(value))
	err := binary.Write(buffer, binary.LittleEndian, length)
	if err != nil {
		return err
	}
	_, err = buffer.Write(value)
	return err
}

func WriteInt16(buffer *bytes.Buffer, value int16) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteInt32(buffer *bytes.Buffer, value int32) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteFloat32(buffer *bytes.Buffer, value float32) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteFloat64(buffer *bytes.Buffer, value float64) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteUInt32(buffer *bytes.Buffer, value uint32) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteUInt16(buffer *bytes.Buffer, value uint16) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteInt64(buffer *bytes.Buffer, value int64) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteByte(buffer *bytes.Buffer, value byte) error {
	return buffer.WriteByte(value)
}

func TimeToInt64(dateTime time.Time) int64 {
	const unixEpochTicks = 621355968000000000
	const nanosecondsToTicks = 100
	ticks := unixEpochTicks + (dateTime.UnixNano() / nanosecondsToTicks)
	return ticks
}

func WriteDateTime(buffer *bytes.Buffer, dateTime time.Time) error {

	const OLEAutomationEpoch = "1899-12-30T00:00:00Z"
	oleEpoch, _ := time.Parse(time.RFC3339, OLEAutomationEpoch)

	duration := dateTime.Sub(oleEpoch)
	fl := float64(duration.Hours() / 24.0)

	return binary.Write(buffer, binary.LittleEndian, fl)
}

func ReadFileAsString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func CompressZlib(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	writer, _ := zlib.NewWriterLevel(&buffer, zlib.BestCompression)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func FileToZlibCompressed(filename string) ([]byte, error) {
	fileContent, err := ReadFileAsString(filename)
	if err != nil {
		return nil, err
	}

	compressedData, err := CompressZlib([]byte(fileContent))
	if err != nil {
		return nil, err
	}

	return compressedData, nil
}

func WritePackedUInt64(buf *bytes.Buffer, value uint64) error {
	var data []byte

	switch {
	case value <= 240:
		data = append(data, byte(value))
	case value <= 2287:
		data = append(data, byte((value-240)/256+241))
		data = append(data, byte((value-240)%256))
	case value <= 67823:
		data = append(data, 249)
		data = append(data, byte((value-2288)/256))
		data = append(data, byte((value-2288)%256))
	case value <= 16777215:
		data = append(data, 250)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
	case value <= 4294967295:
		data = append(data, 251)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
	case value <= 1099511627775:
		data = append(data, 252)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
		data = append(data, byte((value>>32)&255))
	case value <= 281474976710655:
		data = append(data, 253)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
		data = append(data, byte((value>>32)&255))
		data = append(data, byte((value>>40)&255))
	case value <= 72057594037927935:
		data = append(data, 254)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
		data = append(data, byte((value>>32)&255))
		data = append(data, byte((value>>40)&255))
		data = append(data, byte((value>>48)&255))
	default:
		data = append(data, 255)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
		data = append(data, byte((value>>32)&255))
		data = append(data, byte((value>>40)&255))
		data = append(data, byte((value>>48)&255))
		data = append(data, byte((value>>56)&255))
	}

	_, err := buf.Write(data)
	return err
}

func WritePackedUInt32(buf *bytes.Buffer, value uint32) error {
	var data []byte

	if value <= 240 {
		data = append(data, byte(value))
	} else if value <= 2287 {
		data = append(data, byte((value-240)/256+241))
		data = append(data, byte((value-240)%256))
	} else if value <= 67823 {
		data = append(data, 249)
		data = append(data, byte((value-2288)/256))
		data = append(data, byte((value-2288)%256))
	} else if value <= 16777215 {
		data = append(data, 250)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
	} else if value <= 4294967295 {
		data = append(data, 251)
		data = append(data, byte(value&255))
		data = append(data, byte((value>>8)&255))
		data = append(data, byte((value>>16)&255))
		data = append(data, byte((value>>24)&255))
	} else {
		return errors.New("value out of range")
	}

	_, err := buf.Write(data)
	return err
}

func write7BitEncodedInt(buffer *bytes.Buffer, value int) error {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if value != 0 {
			b |= 0x80
		}
		err := buffer.WriteByte(b)
		if err != nil {
			return err
		}
		if value == 0 {
			break
		}
	}
	return nil
}

func WriteUTF16String(buffer *bytes.Buffer, value string) error {
	runeSlice := []rune(value)
	utf16Slice := utf16.Encode(runeSlice)

	err := write7BitEncodedInt(buffer, len(utf16Slice))
	if err != nil {
		return err
	}

	for _, u16 := range utf16Slice {
		err := binary.Write(buffer, binary.LittleEndian, u16)
		if err != nil {
			return err
		}
	}

	return nil
}
