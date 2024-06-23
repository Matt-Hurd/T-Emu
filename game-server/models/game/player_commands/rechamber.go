package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandRechamber struct {
	AmmoToLoad           string
	ShellToUnload        string
	AmmoToUnload         string
	InInventory          bool
	InMisfireMalfunction bool
}

func (msg *PlayerCommandRechamber) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoad); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.ShellToUnload); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.AmmoToUnload); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandRechamber) Deserialize(buffer *bytes.Buffer) error {
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
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}
