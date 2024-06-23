package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandPullOutMagazine struct {
	SlotModeID           string
	AmmoInChamber        int32
	MagTypeCurrent       int32
	Boltcatch            bool
	InMisfireMalfunction bool
}

func (msg *PlayerCommandPullOutMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.SlotModeID); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.AmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.MagTypeCurrent); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.Boltcatch); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandPullOutMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.SlotModeID); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.AmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.MagTypeCurrent); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.Boltcatch); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}
