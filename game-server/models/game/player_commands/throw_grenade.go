package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/enums"
)

type PlayerCommandThrowGrenade struct {
	Variation enums.GrenadeAttackVariation
	Item      core.ComponentialItem
}

func (msg *PlayerCommandThrowGrenade) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, byte(msg.Variation)); err != nil {
		return err
	}
	if err := msg.Item.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandThrowGrenade) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var Variation byte
	if err = helpers.ReadByte(buffer, &Variation); err != nil {
		return err
	}
	msg.Variation = enums.GrenadeAttackVariation(Variation)
	if err = msg.Item.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
