package descriptors

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type RepairEnhancementComponentDescriptor struct {
	BuffType            enums.RepairBuffType
	BuffRarity          enums.BuffRarity
	Value               float32
	ThresholdDurability float32
}

func (repairEnhancementComponent *RepairEnhancementComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteInt32(buffer, int32(repairEnhancementComponent.BuffType))
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(repairEnhancementComponent.BuffRarity))
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, repairEnhancementComponent.Value)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, repairEnhancementComponent.ThresholdDurability)
	if err != nil {
		return err
	}
	return nil
}

func (repairEnhancementComponent *RepairEnhancementComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var val int32
	err = helpers.ReadInt32(buffer, &val)
	repairEnhancementComponent.BuffType = enums.RepairBuffType(val)
	if err != nil {
		return err
	}
	err = helpers.ReadInt32(buffer, &val)
	repairEnhancementComponent.BuffRarity = enums.BuffRarity(val)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &repairEnhancementComponent.Value)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &repairEnhancementComponent.ThresholdDurability)
	if err != nil {
		return err
	}
	return nil
}
