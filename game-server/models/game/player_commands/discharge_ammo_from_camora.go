package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandDischargeAmmoFromCamora struct {
	DischargeAmmo bool
	CamoraIndex   int32
}

func (msg *PlayerCommandDischargeAmmoFromCamora) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.DischargeAmmo); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.CamoraIndex); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDischargeAmmoFromCamora) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.DischargeAmmo); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.CamoraIndex); err != nil {
		return err
	}
	return nil
}
