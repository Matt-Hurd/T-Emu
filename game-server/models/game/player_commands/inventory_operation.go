package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core/descriptors"
)

type PlayerCommandInventoryOperation struct {
	hasOperation bool
	Operation    descriptors.Serializable
}

func (msg *PlayerCommandInventoryOperation) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.hasOperation); err != nil {
		return err
	}
	if msg.Operation == nil {
		return nil
	}

	tmpBuffer := new(bytes.Buffer)
	if err := descriptors.WritePolymorph(buffer, msg.Operation); err != nil {
		return err
	}
	if err := helpers.WriteBytesAndSize(buffer, tmpBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandInventoryOperation) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.hasOperation); err != nil {
		return err
	}
	if !msg.hasOperation {
		return nil
	}

	var tmpBuffer []byte
	if err = helpers.ReadBytesAndSize(buffer, &tmpBuffer); err != nil {
		return err
	}
	err = descriptors.ReadPolymorph(bytes.NewBuffer(tmpBuffer), &msg.Operation)
	if err != nil {
		return err
	}
	return nil
}
