package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type RepairableComponentDescriptor struct {
	Durability    float32
	MaxDurability float32
}

func (repairableComponent *RepairableComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteFloat32(buffer, repairableComponent.Durability)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, repairableComponent.MaxDurability)
	if err != nil {
		return err
	}
	return nil
}

func (repairableComponent *RepairableComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadFloat32(buffer, &repairableComponent.Durability)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &repairableComponent.MaxDurability)
	if err != nil {
		return err
	}
	return nil
}
