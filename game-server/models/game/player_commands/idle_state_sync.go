package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandIdleStateSync struct {
	InInventory  bool
	InLeftStance bool
}

func (msg *PlayerCommandIdleStateSync) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InLeftStance); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandIdleStateSync) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InLeftStance); err != nil {
		return err
	}
	return nil
}
