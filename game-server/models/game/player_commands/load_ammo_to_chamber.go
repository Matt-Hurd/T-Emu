package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandLoadAmmoToChamber struct {
	AmmoTemplate string
	ChamberIndex int32
}

func (msg *PlayerCommandLoadAmmoToChamber) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoTemplate); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ChamberIndex); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandLoadAmmoToChamber) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoTemplate); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ChamberIndex); err != nil {
		return err
	}
	return nil
}
