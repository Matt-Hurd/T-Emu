package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSkillParameters struct {
	CovertMovementSpeed                float32
	CovertMOvementLoud                 float32
	CovertMovementSoundVolume          float32
	CovertMovementEquipment            float32
	BotSoundCoef                       float32
	HeavyVestNoBodyDamageDeflectChance bool
	WeaponDurabilityLosOnShotReduce    float32
	DrawSound                          float32
}

func (msg *PlayerCommandSkillParameters) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, msg.CovertMovementSpeed)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.CovertMOvementLoud)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.CovertMovementSoundVolume)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.CovertMovementEquipment)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.BotSoundCoef)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, msg.HeavyVestNoBodyDamageDeflectChance)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.WeaponDurabilityLosOnShotReduce)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.DrawSound)
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSkillParameters) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &msg.CovertMovementSpeed)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.CovertMOvementLoud)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.CovertMovementSoundVolume)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.CovertMovementEquipment)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.BotSoundCoef)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &msg.HeavyVestNoBodyDamageDeflectChance)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.WeaponDurabilityLosOnShotReduce)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.DrawSound)
	if err != nil {
		return err
	}
	return nil
}
