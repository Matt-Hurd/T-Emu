package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetFinishReloadInternalMagWithOpenBolt struct {
	MagazineCount    int32
	ChamberAmmoCount int32
	InInventory      bool
}

func (msg *PlayerCommandSetFinishReloadInternalMagWithOpenBolt) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.MagazineCount); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ChamberAmmoCount); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetFinishReloadInternalMagWithOpenBolt) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.MagazineCount); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ChamberAmmoCount); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	return nil
}
