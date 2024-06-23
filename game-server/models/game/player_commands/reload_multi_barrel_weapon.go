package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandReloadMultiBarrelWeapon struct {
	AmmoToLoad      []string
	ShellToUnload   []string
	AmmosInChambers int32
	InInventory     bool
}

func (msg *PlayerCommandReloadMultiBarrelWeapon) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, int32(len(msg.AmmoToLoad))); err != nil {
		return err
	}
	for _, ammo := range msg.AmmoToLoad {
		if err := helpers.WriteMongoId(buffer, ammo); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.ShellToUnload))); err != nil {
		return err
	}
	for _, shell := range msg.ShellToUnload {
		if err := helpers.WriteMongoId(buffer, shell); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, msg.AmmosInChambers); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandReloadMultiBarrelWeapon) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var ammoCount int32
	if err = helpers.ReadInt32(buffer, &ammoCount); err != nil {
		return err
	}
	msg.AmmoToLoad = make([]string, ammoCount)
	for i := 0; i < int(ammoCount); i++ {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoad[i]); err != nil {
			return err
		}
	}
	var shellCount int32
	if err = helpers.ReadInt32(buffer, &shellCount); err != nil {
		return err
	}
	msg.ShellToUnload = make([]string, shellCount)
	for i := 0; i < int(shellCount); i++ {
		if err = helpers.ReadMongoId(buffer, &msg.ShellToUnload[i]); err != nil {
			return err
		}
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmosInChambers); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	return nil
}
