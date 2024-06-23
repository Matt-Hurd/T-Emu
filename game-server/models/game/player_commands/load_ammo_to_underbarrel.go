package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandLoadAmmoToUnderbarrel struct {
	AmmoToLoad    string
	ShellToUnload string
}

func (msg *PlayerCommandLoadAmmoToUnderbarrel) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoad); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.ShellToUnload); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandLoadAmmoToUnderbarrel) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoad); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.ShellToUnload); err != nil {
		return err
	}
	return nil
}
