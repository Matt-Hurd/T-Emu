package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetLeftStance struct {
	SetLeftStance bool
}

func (msg *PlayerCommandSetLeftStance) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.SetLeftStance); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetLeftStance) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.SetLeftStance); err != nil {
		return err
	}
	return nil
}
