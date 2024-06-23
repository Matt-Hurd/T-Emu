package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type StationaryCommand byte

const (
	Occupy StationaryCommand = iota
	Leave
	Denied
)

type PlayerCommandSetStationaryWeapon struct {
	StationaryCommand  StationaryCommand
	StationaryWeaponId string
	AmmoInMagCount     int32
	MagazineSlotModeID string
	Magazine           core.ComponentialItem
	ChamberSlotModeID  string
	ChamberAmmo        core.ComponentialItem
}

func (msg *PlayerCommandSetStationaryWeapon) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, byte(msg.StationaryCommand)); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.StationaryWeaponId); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInMagCount); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.MagazineSlotModeID); err != nil {
		return err
	}
	if err := msg.Magazine.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.ChamberSlotModeID); err != nil {
		return err
	}
	if err := msg.ChamberAmmo.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetStationaryWeapon) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.StationaryCommand)); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.StationaryWeaponId); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInMagCount); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.MagazineSlotModeID); err != nil {
		return err
	}
	if err = msg.Magazine.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.ChamberSlotModeID); err != nil {
		return err
	}
	if err = msg.ChamberAmmo.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
