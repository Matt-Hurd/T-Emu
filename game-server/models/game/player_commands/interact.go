package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandInteract struct {
	IsInteract     bool
	AnimationIndex int32
}

func (msg *PlayerCommandInteract) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.IsInteract); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AnimationIndex); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandInteract) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.IsInteract); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AnimationIndex); err != nil {
		return err
	}
	return nil
}
