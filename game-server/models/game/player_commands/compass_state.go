package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandCompassState struct {
	IsActive bool
}

func (msg *PlayerCommandCompassState) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.IsActive); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandCompassState) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.IsActive); err != nil {
		return err
	}
	return nil
}
