package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandDischargeChamber struct {
	AmmoTemplate         string
	ChamberIndex         int32
	InMisfireMalfunction bool
}

func (msg *PlayerCommandDischargeChamber) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, msg.AmmoTemplate); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ChamberIndex); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDischargeChamber) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &msg.AmmoTemplate); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ChamberIndex); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}
