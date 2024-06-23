package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type ArmorInfo struct {
	ItemID                 string
	ArmorType              enums.ArmorType
	MaxDurability          float32
	Durability             float32
	TemplateDurability     int32
	RicochetValues         Vector3
	ArmorClass             int32
	Material               enums.MaterialType
	ArmorColliders         []enums.BodyPartColliderType
	ArmorPlateColliderMask enums.ArmorPlateCollider
	IsComposite            bool
	IsToggledAndOff        bool
}

func (a *ArmorInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &a.ItemID); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&a.ArmorType)); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &a.MaxDurability); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &a.Durability); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &a.TemplateDurability); err != nil {
		return err
	}
	if err = a.RicochetValues.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &a.ArmorClass); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&a.Material)); err != nil {
		return err
	}
	var cnt int32
	if err = helpers.ReadInt32(buffer, &cnt); err != nil {
		return err
	}
	a.ArmorColliders = make([]enums.BodyPartColliderType, cnt)
	var tmp byte
	for i := 0; i < int(cnt); i++ {
		if err = helpers.ReadByte(buffer, &tmp); err != nil {
			return err
		}
		a.ArmorColliders[i] = enums.BodyPartColliderType(tmp)
	}
	var armorPlateColliderMask uint16
	if err = helpers.ReadUInt16(buffer, &armorPlateColliderMask); err != nil {
		return err
	}
	a.ArmorPlateColliderMask = enums.ArmorPlateCollider(armorPlateColliderMask)
	if err = helpers.ReadBool(buffer, &a.IsComposite); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &a.IsToggledAndOff); err != nil {
		return err
	}
	return nil
}

func (a *ArmorInfo) Serialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.WriteString(buffer, a.ItemID); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(a.ArmorType)); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, a.MaxDurability); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, a.Durability); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, a.TemplateDurability); err != nil {
		return err
	}
	if err = a.RicochetValues.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, a.ArmorClass); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(a.Material)); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, int32(len(a.ArmorColliders))); err != nil {
		return err
	}
	for _, v := range a.ArmorColliders {
		if err = helpers.WriteByte(buffer, byte(v)); err != nil {
			return err
		}
	}
	if err = helpers.WriteUInt16(buffer, uint16(a.ArmorPlateColliderMask)); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, a.IsComposite); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, a.IsToggledAndOff); err != nil {
		return err
	}
	return nil
}
