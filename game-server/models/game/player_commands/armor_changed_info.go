package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/enums"
)

type PlayerCommandArmorChangedInfo struct {
	RemoveItemIds          []string
	Added                  bool
	ArmorPlateColliderMask enums.ArmorPlateCollider
	ArmorInfos             []core.ArmorInfo
}

func (msg *PlayerCommandArmorChangedInfo) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, int32(len(msg.RemoveItemIds))); err != nil {
		return err
	}
	for _, v := range msg.RemoveItemIds {
		if err := helpers.WriteString(buffer, v); err != nil {
			return err
		}
	}
	if err := helpers.WriteBool(buffer, msg.Added); err != nil {
		return err
	}
	if err := helpers.WriteUInt16(buffer, uint16(msg.ArmorPlateColliderMask)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, int32(len(msg.ArmorInfos))); err != nil {
		return err
	}
	for _, v := range msg.ArmorInfos {
		if err := v.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandArmorChangedInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var removeItemIdCount int32
	if err = helpers.ReadInt32(buffer, &removeItemIdCount); err != nil {
		return err
	}
	msg.RemoveItemIds = make([]string, removeItemIdCount)
	for i := 0; i < int(removeItemIdCount); i++ {
		if err = helpers.ReadString(buffer, &msg.RemoveItemIds[i]); err != nil {
			return err
		}
	}
	if err = helpers.ReadBool(buffer, &msg.Added); err != nil {
		return err
	}
	var armorPlateColliderMask uint16
	if err = helpers.ReadUInt16(buffer, &armorPlateColliderMask); err != nil {
		return err
	}
	msg.ArmorPlateColliderMask = enums.ArmorPlateCollider(armorPlateColliderMask)
	var armorInfoCount int32
	if err = helpers.ReadInt32(buffer, &armorInfoCount); err != nil {
		return err
	}
	msg.ArmorInfos = make([]core.ArmorInfo, armorInfoCount)
	for i := 0; i < int(armorInfoCount); i++ {
		if err = msg.ArmorInfos[i].Deserialize(buffer); err != nil {
			return err
		}
	}
	return nil
}
