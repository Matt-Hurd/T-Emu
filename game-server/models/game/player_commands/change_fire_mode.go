package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandChangeFireMode struct {
	FireMode enums.FireMode
}

func (msg *PlayerCommandChangeFireMode) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteByte(buffer, byte(msg.FireMode))
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandChangeFireMode) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var fireMode byte
	if err = helpers.ReadByte(buffer, &fireMode); err != nil {
		return err
	}
	msg.FireMode = enums.FireMode(fireMode)
	return nil
}
