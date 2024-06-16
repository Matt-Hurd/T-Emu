package helpers

import (
	"bytes"
	"encoding/binary"
	"errors"
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
