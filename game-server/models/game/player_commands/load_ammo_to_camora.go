package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandLoadAmmoToCamora struct {
	LoadAmmo    bool
	RemoveShell bool
	CamoraIndex int32
	AmmoToLoad  string
}

func (msg *PlayerCommandLoadAmmoToCamora) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.LoadAmmo); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.RemoveShell); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.CamoraIndex); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoad); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandLoadAmmoToCamora) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.LoadAmmo); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.RemoveShell); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.CamoraIndex); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoad); err != nil {
		return err
	}
	return nil
}
