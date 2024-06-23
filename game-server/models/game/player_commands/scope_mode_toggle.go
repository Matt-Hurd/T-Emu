package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandScopeModeToggle struct {
	ID                    string
	SetSilently           bool
	ScopeMode             int32
	ScopeIndexInsideSight int32
	ScopeCalibrationIndex int32
}

func (msg *PlayerCommandScopeModeToggle) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.ID); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.SetSilently); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ScopeMode); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ScopeIndexInsideSight); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ScopeCalibrationIndex); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandScopeModeToggle) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.ID); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.SetSilently); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ScopeMode); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ScopeIndexInsideSight); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ScopeCalibrationIndex); err != nil {
		return err
	}
	return nil
}
