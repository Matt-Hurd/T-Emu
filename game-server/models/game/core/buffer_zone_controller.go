package core

import (
	"bytes"
	"game-server/helpers"
)

type BufferZoneControllerClass struct {
	IsAvailable bool
}

func (b *BufferZoneControllerClass) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadBool(buffer, &b.IsAvailable); err != nil {
		return err
	}
	return nil
}

func (b *BufferZoneControllerClass) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteBool(buffer, b.IsAvailable); err != nil {
		return err
	}
	return nil
}
