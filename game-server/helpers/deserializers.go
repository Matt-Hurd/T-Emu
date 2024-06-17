package helpers

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
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

func ReadInt32(buffer *bytes.Buffer, result *int32) error {
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

func ReadDateTime(buffer *bytes.Buffer, result *time.Time) error {
	var binaryDate int64
	err := binary.Read(buffer, binary.LittleEndian, &binaryDate)
	if err != nil {
		return err
	}
	const ticksToNanoseconds = 100
	const unixEpochTicks = 621355968000000000
	ticks := binaryDate & 0x3FFFFFFFFFFFFFFF

	nanoseconds := (ticks - unixEpochTicks) * ticksToNanoseconds
	gameDateTime := time.Unix(0, nanoseconds).UTC()
	*result = gameDateTime
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
