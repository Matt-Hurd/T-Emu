package core

import (
	"bytes"
	"game-server/helpers"
)

type BTR struct {
	HasBTR  bool
	BTRData []byte
}

func (b *BTR) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadBool(buffer, &b.HasBTR); err != nil {
		return err
	}
	if b.HasBTR {
		if err := helpers.ReadBytesAndSize(buffer, &b.BTRData); err != nil {
			return err
		}
	}
	return nil
}

func (b *BTR) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteBool(buffer, b.HasBTR); err != nil {
		return err
	}
	if b.HasBTR {
		if err = helpers.WriteBytesAndSize(buffer, b.BTRData); err != nil {
			return err
		}
	}
	return nil
}
