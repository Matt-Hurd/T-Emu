package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PlayerCommandInsertMagazine struct {
	InMisfireMalfunction   bool
	MagTypeCurrent         int32
	MagTypeNew             int32
	AmmoInChamberCurrent   int32
	AmmoInChamberResult    int32
	AmmoInMag              int32
	NeedToAddAmmoInChamber bool
	ShellsInWeapon         int32
	WeaponLevel            int32
	ForceEmptyChamber      bool
	Boltcatch              bool
	WeaponItemID           string
	SlotModeID             string
	Items                  core.ComponentialItem
}

func (msg *PlayerCommandInsertMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.InMisfireMalfunction); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.MagTypeCurrent); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.MagTypeNew); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamberCurrent); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamberResult); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInMag); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.NeedToAddAmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ShellsInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.WeaponLevel); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.ForceEmptyChamber); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Boltcatch); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.WeaponItemID); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.SlotModeID); err != nil {
		return err
	}
	if err := msg.Items.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandInsertMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.InMisfireMalfunction); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.MagTypeCurrent); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.MagTypeNew); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamberCurrent); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamberResult); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInMag); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.NeedToAddAmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ShellsInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.WeaponLevel); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.ForceEmptyChamber); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Boltcatch); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.WeaponItemID); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.SlotModeID); err != nil {
		return err
	}
	if err = msg.Items.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
