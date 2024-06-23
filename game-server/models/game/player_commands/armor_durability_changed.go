package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandArmorDurabilityChanged struct {
	ArmorItemID            string
	Durability             float32
	ArmorPlateColliderMask enums.ArmorPlateCollider
}

func (msg *PlayerCommandArmorDurabilityChanged) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.ArmorItemID); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.Durability); err != nil {
		return err
	}
	if err := helpers.WriteUInt16(buffer, uint16(msg.ArmorPlateColliderMask)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandArmorDurabilityChanged) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.ArmorItemID); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.Durability); err != nil {
		return err
	}
	var armorPlateColliderMask uint16
	if err = helpers.ReadUInt16(buffer, &armorPlateColliderMask); err != nil {
		return err
	}
	msg.ArmorPlateColliderMask = enums.ArmorPlateCollider(armorPlateColliderMask)
	return nil
}
