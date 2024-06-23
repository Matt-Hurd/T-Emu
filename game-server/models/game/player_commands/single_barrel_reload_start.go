package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSingleBarrelReloadStart struct {
	AmmoToLoad    string
	ShellToUnload string
	AmmoToUnload  string
	WeaponLevel   int32
	InInventory   bool
}

func (msg *PlayerCommandSingleBarrelReloadStart) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoad); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.ShellToUnload); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.AmmoToUnload); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.WeaponLevel); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSingleBarrelReloadStart) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoad); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.ShellToUnload); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToUnload); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.WeaponLevel); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	return nil
}
