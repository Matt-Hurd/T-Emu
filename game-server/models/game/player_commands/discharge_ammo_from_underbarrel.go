package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandDischargeAmmoFromUnderbarrel struct {
	AmmoTemplate string
}

func (msg *PlayerCommandDischargeAmmoFromUnderbarrel) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDischargeAmmoFromUnderbarrel) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}
