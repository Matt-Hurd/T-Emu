package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandMeleeAttack struct {
	Fire             bool
	AltFire          bool
	QuickFire        bool
	MeleeAttackSpeed float32
}

func (msg *PlayerCommandMeleeAttack) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.Fire); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.AltFire); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.QuickFire); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.MeleeAttackSpeed); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandMeleeAttack) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.Fire); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.AltFire); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.QuickFire); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.MeleeAttackSpeed); err != nil {
		return err
	}
	return nil
}
