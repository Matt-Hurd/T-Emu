package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/enums"
)

type OperationType byte

const (
	OperationTypeEquip OperationType = iota
	OperationTypeUnequip
	OperationTypeToggle
	OperationTypeToggleFold
)

type PlayerCommandEquipChanged struct {
	SlotType      enums.EquipmentSlot
	OperationType OperationType
	IsInInventory bool
	ItemsForEquip core.ComponentialItem
}

func (msg *PlayerCommandEquipChanged) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.SlotType)); err != nil {
		return err
	}
	if err := buffer.WriteByte(byte(msg.OperationType)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.IsInInventory); err != nil {
		return err
	}
	if err := msg.ItemsForEquip.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandEquipChanged) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.SlotType)); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&msg.OperationType)); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.IsInInventory); err != nil {
		return err
	}
	if err = msg.ItemsForEquip.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
