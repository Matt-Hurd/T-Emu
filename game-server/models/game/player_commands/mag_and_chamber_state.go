package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandMagAndChamberState struct {
	AmmoInMag     int32
	AmmoInChamber int32
	MagInWeapon   bool
	Armed         bool
}

func (msg *PlayerCommandMagAndChamberState) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.AmmoInMag); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.MagInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Armed); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandMagAndChamberState) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.AmmoInMag); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.MagInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Armed); err != nil {
		return err
	}
	return nil
}
