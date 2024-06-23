package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandReloadInternalMagazine struct {
	AmmoInChamber    int32
	ShellsInWeapon   int32
	WeaponLevel      int32
	InInventory      bool
	MalfunctionState enums.MalfunctionState
	AmmoInStack      int32
}

func (msg *PlayerCommandReloadInternalMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ShellsInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.WeaponLevel); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(msg.MalfunctionState)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInStack); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandReloadInternalMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ShellsInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.WeaponLevel); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&msg.MalfunctionState)); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInStack); err != nil {
		return err
	}
	return nil
}
