package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandReloadInternalMagWithOpenBolt struct {
	AmmoToLoadTemplate   string
	InInventory          bool
	MagInWeapon          bool
	IsExternalMag        bool
	AmmoPackCount        int32
	MagFreeSpace         int32
	InMisfireMalfunction bool
	HasAmmoInChamber     bool
}

func (msg *PlayerCommandReloadInternalMagWithOpenBolt) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoadTemplate); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.MagInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.IsExternalMag); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoPackCount); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.MagFreeSpace); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InMisfireMalfunction); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.HasAmmoInChamber); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandReloadInternalMagWithOpenBolt) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoadTemplate); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.MagInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.IsExternalMag); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoPackCount); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.MagFreeSpace); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InMisfireMalfunction); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.HasAmmoInChamber); err != nil {
		return err
	}
	return nil
}
