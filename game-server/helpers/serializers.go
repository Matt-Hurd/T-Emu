package helpers

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"os"
	"time"
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

func WriteInt32(buffer *bytes.Buffer, value int32) error {
	return binary.Write(buffer, binary.LittleEndian, value)
}

func WriteFloat32(buffer *bytes.Buffer, value float32) error {
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
