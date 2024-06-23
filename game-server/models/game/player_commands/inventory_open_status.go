package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandInventoryOpenStatus struct {
	InventoryOpen bool
}

func (msg *PlayerCommandInventoryOpenStatus) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.InventoryOpen); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandInventoryOpenStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.InventoryOpen); err != nil {
		return err
	}
	return nil
}
