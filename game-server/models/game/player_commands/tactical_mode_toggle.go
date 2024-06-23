package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandTacticalModeToggle struct {
	ID          string
	SetSilently bool
	State       bool
	LightMode   int32
}

func (msg *PlayerCommandTacticalModeToggle) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.ID); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.SetSilently); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.State); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.LightMode); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandTacticalModeToggle) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.ID); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.SetSilently); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.State); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.LightMode); err != nil {
		return err
	}
	return nil
}
