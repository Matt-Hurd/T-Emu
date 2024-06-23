package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandCylinderCamoraIndex struct {
	CamoraIndex  int32
	HammerClosed bool
}

func (msg *PlayerCommandCylinderCamoraIndex) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.CamoraIndex); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.HammerClosed); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandCylinderCamoraIndex) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.CamoraIndex); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.HammerClosed); err != nil {
		return err
	}
	return nil
}
