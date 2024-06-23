package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandReloadCylinderMagazine struct {
	AmmoForLoad        string
	ReloadFast         bool
	CurrentCamoraIndex int32
	WeaponLevel        int32
	InInventory        bool
	AmmoInMag          int32
	FreeCamoras        []int32
	Shells             []int32
}

func (msg *PlayerCommandReloadCylinderMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoForLoad); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.ReloadFast); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.CurrentCamoraIndex); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.WeaponLevel); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInMag); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.FreeCamoras))); err != nil {
		return err
	}
	for _, camora := range msg.FreeCamoras {
		if err := helpers.WriteInt32(buffer, camora); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.Shells))); err != nil {
		return err
	}
	for _, shell := range msg.Shells {
		if err := helpers.WriteInt32(buffer, shell); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandReloadCylinderMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoForLoad); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.ReloadFast); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.CurrentCamoraIndex); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.WeaponLevel); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInMag); err != nil {
		return err
	}
	var length int32
	if err = helpers.ReadInt32(buffer, &length); err != nil {
		return err
	}
	msg.FreeCamoras = make([]int32, length)
	for i := range msg.FreeCamoras {
		if err = helpers.ReadInt32(buffer, &msg.FreeCamoras[i]); err != nil {
			return err
		}
	}
	if err = helpers.ReadInt32(buffer, &length); err != nil {
		return err
	}
	msg.Shells = make([]int32, length)
	for i := range msg.Shells {
		if err = helpers.ReadInt32(buffer, &msg.Shells[i]); err != nil {
			return err
		}
	}
	return nil
}
