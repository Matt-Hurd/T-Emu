package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandBtrGoInInteraction struct {
	BtrSideId byte
	BtrSlotId byte
	Fast      bool
}

func (msg *PlayerCommandBtrGoInInteraction) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, msg.BtrSideId); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, msg.BtrSlotId); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Fast); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandBtrGoInInteraction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, &msg.BtrSideId); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &msg.BtrSlotId); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Fast); err != nil {
		return err
	}
	return nil
}
