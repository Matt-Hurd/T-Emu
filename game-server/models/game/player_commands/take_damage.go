package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandTakeDamage struct {
	Damage     float32
	DamageType enums.DamageType
}

func (msg *PlayerCommandTakeDamage) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteFloat32(buffer, msg.Damage); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, int32(msg.DamageType)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandTakeDamage) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadFloat32(buffer, &msg.Damage); err != nil {
		return err
	}
	var DamageType int32
	if err = helpers.ReadInt32(buffer, &DamageType); err != nil {
		return err
	}
	msg.DamageType = enums.DamageType(DamageType)
	return nil
}
