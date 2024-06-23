package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandSwingGrenade struct {
	Variation enums.GrenadeAttackVariation
}

func (msg *PlayerCommandSwingGrenade) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.Variation)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSwingGrenade) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.Variation)); err != nil {
		return err
	}
	return nil
}
