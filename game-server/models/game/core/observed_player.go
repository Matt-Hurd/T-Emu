package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core/descriptors"
	"game-server/models/game/enums"
	"game-server/models/game/math"
)

type ObservedPlayer struct {
	RemoteTime             float32
	BodyPosition           math.Vector3
	Customization          map[enums.BodyModelPart]string
	Side                   enums.PlayerSide
	WildSpawnType          enums.WildSpawnType
	GroupID                string
	TeamID                 string
	IsAI                   bool
	ProfileID              string
	Voice                  string
	VoIPState              enums.VoipState
	NickName               string
	AccountId              string
	ArmorPlateColliderMask enums.ArmorPlateCollider
	InventoryDescriptor    descriptors.InventoryDescriptor // Should technically be deserialized into proper inventory
	HandsController        descriptors.Serializable        // PlayerCommandSetHands (should technically be deserialized into proper controller)
	ArmorsInfo             []ArmorInfo
	// ArenaObservedPlayerSpawnMessage ArenaObservedPlayerSpawnMessage
}

func (p *ObservedPlayer) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteFloat32(buffer, p.RemoteTime); err != nil {
		return err
	}
	if err := p.BodyPosition.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(len(p.Customization))); err != nil {
		return err
	}
	for k, v := range p.Customization {
		if err := helpers.WriteByte(buffer, byte(k)); err != nil {
			return err
		}
		if err := helpers.WriteMongoId(buffer, v); err != nil {
			return err
		}
	}
	if err := helpers.WriteByte(buffer, byte(p.Side)); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(p.WildSpawnType)); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.GroupID); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.TeamID); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.IsAI); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.ProfileID); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.Voice); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(p.VoIPState)); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.NickName); err != nil {
		return err
	}
	if err := helpers.WriteStringPlus(buffer, p.AccountId); err != nil {
		return err
	}
	if err := helpers.WriteUInt16(buffer, uint16(p.ArmorPlateColliderMask)); err != nil {
		return err
	}
	if err := p.InventoryDescriptor.Serialize(buffer); err != nil {
		return err
	}
	if err := p.HandsController.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(len(p.ArmorsInfo))); err != nil {
		return err
	}
	for _, v := range p.ArmorsInfo {
		if err := v.Serialize(buffer); err != nil {
			return err
		}
	}
	// if err := p.ArenaObservedPlayerSpawnMessage.Serialize(buffer); err != nil {
	// 	return err
	// }
	return nil
}

func (p *ObservedPlayer) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadFloat32(buffer, &p.RemoteTime); err != nil {
		return err
	}
	if err = p.BodyPosition.Deserialize(buffer); err != nil {
		return err
	}
	var count byte
	if err = helpers.ReadByte(buffer, &count); err != nil {
		return err
	}
	p.Customization = make(map[enums.BodyModelPart]string)
	for i := 0; i < int(count); i++ {
		var key byte
		if err = helpers.ReadByte(buffer, &key); err != nil {
			return err
		}
		var value string
		if err = helpers.ReadMongoId(buffer, &value); err != nil {
			return err
		}
		p.Customization[enums.BodyModelPart(key)] = value
	}
	if err = helpers.ReadByte(buffer, (*byte)(&p.Side)); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&p.WildSpawnType)); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.GroupID); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.TeamID); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.IsAI); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.ProfileID); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.Voice); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&p.VoIPState)); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.NickName); err != nil {
		return err
	}
	if err = helpers.ReadStringMinus(buffer, &p.AccountId); err != nil {
		return err
	}
	var armorPlateColliderMask uint16
	if err = helpers.ReadUInt16(buffer, &armorPlateColliderMask); err != nil {
		return err
	}
	p.ArmorPlateColliderMask = enums.ArmorPlateCollider(armorPlateColliderMask)
	if err = p.InventoryDescriptor.Deserialize(buffer); err != nil {
		return err
	}
	if err = p.HandsController.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &count); err != nil {
		return err
	}
	p.ArmorsInfo = make([]ArmorInfo, count)
	for i := 0; i < int(count); i++ {
		if err = p.ArmorsInfo[i].Deserialize(buffer); err != nil {
			return err
		}
	}
	// if err = p.ArenaObservedPlayerSpawnMessage.Deserialize(buffer); err != nil {
	// 	return err
	// }
	return nil
}
