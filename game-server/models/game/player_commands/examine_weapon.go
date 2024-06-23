package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandExamineWeapon struct {
	CheckAmmo        bool
	CheckChamber     bool
	Look             bool
	MalfunctionState enums.MalfunctionState
	CheckFireMode    bool
	InInventory      bool
}

func (msg *PlayerCommandExamineWeapon) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.CheckAmmo); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.CheckChamber); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Look); err != nil {
		return err
	}
	if err := buffer.WriteByte(byte(msg.MalfunctionState)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.CheckFireMode); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InInventory); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandExamineWeapon) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.CheckAmmo); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.CheckChamber); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Look); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&msg.MalfunctionState)); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.CheckFireMode); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InInventory); err != nil {
		return err
	}
	return nil
}
