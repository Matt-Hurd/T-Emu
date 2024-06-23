package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetExternalMagazineState struct {
	CanReload      bool
	Active         bool
	ResetLeftHand  bool
	ShellsInWeapon int32
	InInventory    bool
	MagInWeapon    bool
}

func (msg *PlayerCommandSetExternalMagazineState) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.CanReload); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Active); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.ResetLeftHand); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ShellsInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.MagInWeapon); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetExternalMagazineState) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.CanReload); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Active); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.ResetLeftHand); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ShellsInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.MagInWeapon); err != nil {
		return err
	}
	return nil
}
