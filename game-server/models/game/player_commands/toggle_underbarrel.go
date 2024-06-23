package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandToggleUnderbarrel struct {
	ToggleOn bool
}

func (msg *PlayerCommandToggleUnderbarrel) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.ToggleOn); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandToggleUnderbarrel) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.ToggleOn); err != nil {
		return err
	}
	return nil
}
