package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandUnderbarrelReload struct {
	AmmoToLoad    string
	ShellToUnload string
}

func (msg *PlayerCommandUnderbarrelReload) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoToLoad); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.ShellToUnload); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandUnderbarrelReload) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoad); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.ShellToUnload); err != nil {
		return err
	}
	return nil
}
