package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSyncCylinderMagazine struct {
	CurrentCamoraIndex int32
	AmmoInChambers     []string
	ShellsInChambers   []string
}

func (msg *PlayerCommandSyncCylinderMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.CurrentCamoraIndex); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.AmmoInChambers))); err != nil {
		return err
	}
	for _, value := range msg.AmmoInChambers {
		if err := helpers.WriteMongoId(buffer, value); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.ShellsInChambers))); err != nil {
		return err
	}
	for _, value := range msg.ShellsInChambers {
		if err := helpers.WriteMongoId(buffer, value); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandSyncCylinderMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.CurrentCamoraIndex); err != nil {
		return err
	}
	var AmmoInChambersLength int32
	if err = helpers.ReadInt32(buffer, &AmmoInChambersLength); err != nil {
		return err
	}
	msg.AmmoInChambers = make([]string, AmmoInChambersLength)
	for i := range msg.AmmoInChambers {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoInChambers[i]); err != nil {
			return err
		}
	}
	var ShellsInChambersLength int32
	if err = helpers.ReadInt32(buffer, &ShellsInChambersLength); err != nil {
		return err
	}
	msg.ShellsInChambers = make([]string, ShellsInChambersLength)
	for i := range msg.ShellsInChambers {
		if err = helpers.ReadMongoId(buffer, &msg.ShellsInChambers[i]); err != nil {
			return err
		}
	}
	return nil
}
