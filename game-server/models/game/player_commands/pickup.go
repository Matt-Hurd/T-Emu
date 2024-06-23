package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandPickup struct {
	Pickup bool
}

func (msg *PlayerCommandPickup) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.Pickup); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandPickup) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.Pickup); err != nil {
		return err
	}
	return nil
}
