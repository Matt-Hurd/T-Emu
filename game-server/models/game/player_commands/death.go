package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/core/descriptors"
	"game-server/models/game/enums"
)

type PlayerCommandDeath struct {
	CorpseImpulse       core.CorpseImpulse
	LastDamagedBodyPart enums.BodyPart
	LastDamageType      enums.DamageType
	Inventory           descriptors.InventoryDescriptor
}

func (msg *PlayerCommandDeath) Serialize(buffer *bytes.Buffer) error {
	err := msg.CorpseImpulse.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, byte(msg.LastDamagedBodyPart))
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, byte(msg.LastDamageType))
	if err != nil {
		return err
	}
	err = msg.Inventory.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDeath) Deserialize(buffer *bytes.Buffer) error {
	var val byte
	err := msg.CorpseImpulse.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.ReadByte(buffer, &val)
	msg.LastDamagedBodyPart = enums.BodyPart(val)
	if err != nil {
		return err
	}
	err = helpers.ReadByte(buffer, &val)
	msg.LastDamageType = enums.DamageType(val)
	if err != nil {
		return err
	}
	err = msg.Inventory.Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}
