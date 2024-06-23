package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1510
type ItemDescriptor struct {
	Id                        string
	TemplateId                string
	StackCount                int32
	SpawnedInSession          bool
	ActiveCamora              byte
	IsUnderBarrelDeviceActive bool
	Components                []ItemComponent
	Slots                     []SlotDescriptor
	ShellsInWeapon            []string
	ShellsInUnderbarrelWeapon []string
	Grids                     []LootGridDescriptor
	StackSlots                []StackSlotDescriptor
	Malfunction               []MalfunctionDescriptor
}

// GClass1511
type ItemComponent interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

func (item *ItemDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, item.Id)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, item.TemplateId)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, item.StackCount)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, item.SpawnedInSession)
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, item.ActiveCamora)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, item.IsUnderBarrelDeviceActive)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(len(item.Components)))
	if err != nil {
		return err
	}
	for _, v := range item.Components {
		err = WritePolymorph(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.Slots)))
	if err != nil {
		return err
	}
	for _, v := range item.Slots {
		err = v.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.ShellsInWeapon)))
	if err != nil {
		return err
	}
	for _, v := range item.ShellsInWeapon {
		err = helpers.WriteUTF16String(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.ShellsInUnderbarrelWeapon)))
	if err != nil {
		return err
	}
	for _, v := range item.ShellsInUnderbarrelWeapon {
		err = helpers.WriteUTF16String(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.Grids)))
	if err != nil {
		return err
	}
	for _, v := range item.Grids {
		err = v.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.StackSlots)))
	if err != nil {
		return err
	}
	for _, v := range item.StackSlots {
		err = v.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(item.Malfunction)))
	if err != nil {
		return err
	}
	for _, v := range item.Malfunction {
		err = v.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (item *ItemDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &item.Id)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &item.TemplateId)
	if err != nil {
		return err
	}
	err = helpers.ReadInt32(buffer, &item.StackCount)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &item.SpawnedInSession)
	if err != nil {
		return err
	}
	err = helpers.ReadByte(buffer, &item.ActiveCamora)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &item.IsUnderBarrelDeviceActive)
	if err != nil {
		return err
	}
	var ComponentsLength int32
	err = helpers.ReadInt32(buffer, &ComponentsLength)
	if err != nil {
		return err
	}
	item.Components = make([]ItemComponent, ComponentsLength)
	for i := range item.Components {
		var component Serializable
		err = ReadPolymorph(buffer, &component)
		if err != nil {
			return err
		}
		item.Components[i] = component
	}
	var SlotsLength int32
	err = helpers.ReadInt32(buffer, &SlotsLength)
	if err != nil {
		return err
	}
	item.Slots = make([]SlotDescriptor, SlotsLength)
	for i := range item.Slots {
		err = item.Slots[i].Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var ShellsInWeaponLength int32
	err = helpers.ReadInt32(buffer, &ShellsInWeaponLength)
	if err != nil {
		return err
	}
	item.ShellsInWeapon = make([]string, ShellsInWeaponLength)
	for i := range item.ShellsInWeapon {
		err = helpers.ReadUTF16String(buffer, &item.ShellsInWeapon[i])
		if err != nil {
			return err
		}
	}
	var ShellsInUnderbarrelWeaponLength int32
	err = helpers.ReadInt32(buffer, &ShellsInUnderbarrelWeaponLength)
	if err != nil {
		return err
	}
	item.ShellsInUnderbarrelWeapon = make([]string, ShellsInUnderbarrelWeaponLength)
	for i := range item.ShellsInUnderbarrelWeapon {
		err = helpers.ReadUTF16String(buffer, &item.ShellsInUnderbarrelWeapon[i])
		if err != nil {
			return err
		}
	}
	var GridsLength int32
	err = helpers.ReadInt32(buffer, &GridsLength)
	if err != nil {
		return err
	}
	item.Grids = make([]LootGridDescriptor, GridsLength)
	for i := range item.Grids {
		err = item.Grids[i].Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var StackSlotsLength int32
	err = helpers.ReadInt32(buffer, &StackSlotsLength)
	if err != nil {
		return err
	}
	item.StackSlots = make([]StackSlotDescriptor, StackSlotsLength)
	for i := range item.StackSlots {
		err = item.StackSlots[i].Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var MalfunctionLength int32
	err = helpers.ReadInt32(buffer, &MalfunctionLength)
	if err != nil {
		return err
	}
	item.Malfunction = make([]MalfunctionDescriptor, MalfunctionLength)
	for i := range item.Malfunction {
		err = item.Malfunction[i].Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
