package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/enums"
)

type PlayerCommandSetHands struct {
	HandControllerType           enums.HandsControllerType
	FastHide                     bool
	Armed                        bool
	MalfunctionState             enums.MalfunctionState
	AmmoInChamber                byte
	DrawAnimationSpeedMultiplier float32
	MalfunctionedAmmo            string
	Item                         core.ComponentialItem
}

func (msg *PlayerCommandSetHands) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.HandControllerType)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.FastHide); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Armed); err != nil {
		return err
	}
	if err := buffer.WriteByte(byte(msg.MalfunctionState)); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, msg.AmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.DrawAnimationSpeedMultiplier); err != nil {
		return err
	}
	if msg.MalfunctionState != enums.MalfunctionStateNone {
		if err := helpers.WriteStringPlus(buffer, msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	if err := msg.Item.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetHands) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.HandControllerType)); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.FastHide); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Armed); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&msg.MalfunctionState)); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &msg.AmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.DrawAnimationSpeedMultiplier); err != nil {
		return err
	}
	if msg.MalfunctionState != enums.MalfunctionStateNone {
		if err = helpers.ReadStringMinus(buffer, &msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	if err = msg.Item.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
