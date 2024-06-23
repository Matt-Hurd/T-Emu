package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandRepairMalfunction struct {
	MalfunctionState      enums.MalfunctionState
	AmmoToFire            string
	MalfunctionedAmmo     string
	AmmoToLoadAfterRepair string
	AmmoInMag             int32
}

func (msg *PlayerCommandRepairMalfunction) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, byte(msg.MalfunctionState)); err != nil {
		return err
	}
	flag := msg.MalfunctionedAmmo != ""
	if err := helpers.WriteBool(buffer, flag); err != nil {
		return err
	}
	if flag {
		if err := helpers.WriteMongoId(buffer, msg.AmmoToFire); err != nil {
			return err
		}
	}
	flag2 := msg.MalfunctionedAmmo != "" // TODO: Check if this is only a client error, this is what the client has
	if err := helpers.WriteBool(buffer, flag2); err != nil {
		return err
	}
	if flag2 {
		if err := helpers.WriteMongoId(buffer, msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	flag3 := msg.AmmoToLoadAfterRepair != ""
	if err := helpers.WriteBool(buffer, flag3); err != nil {
		return err
	}
	if flag3 {
		if err := helpers.WriteMongoId(buffer, msg.AmmoToLoadAfterRepair); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInMag); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandRepairMalfunction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.MalfunctionState)); err != nil {
		return err
	}
	var flag bool
	if err = helpers.ReadBool(buffer, &flag); err != nil {
		return err
	}
	if flag {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoToFire); err != nil {
			return err
		}
	}
	var flag2 bool
	if err = helpers.ReadBool(buffer, &flag2); err != nil {
		return err
	}
	if flag2 {
		if err = helpers.ReadMongoId(buffer, &msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	var flag3 bool
	if err = helpers.ReadBool(buffer, &flag3); err != nil {
		return err
	}
	if flag3 {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoToLoadAfterRepair); err != nil {
			return err
		}
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInMag); err != nil {
		return err
	}
	return nil
}
