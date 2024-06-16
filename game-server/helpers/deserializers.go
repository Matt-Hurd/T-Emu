package helpers

import (
	"bytes"
	"encoding/binary"
	"errors"
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
	err := binary.Read(buffer, binary.LittleEndian, result)
	if err != nil {
		return err
	}
	return nil
}
