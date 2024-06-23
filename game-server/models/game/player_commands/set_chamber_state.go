package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetChamberState struct {
	CanReload      bool
	Active         bool
	ResetLeftHand  bool
	AmmoInChamber  int32
	ShellsInWeapon int32
	InInventory    bool
}

func (msg *PlayerCommandSetChamberState) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.CanReload); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Active); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.ResetLeftHand); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ShellsInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetChamberState) Deserialize(buffer *bytes.Buffer) error {
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
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ShellsInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	return nil
}
