package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandAiming struct {
	IsAiming    bool
	AimingIndex int
}

func (msg *PlayerCommandAiming) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteBool(buffer, msg.IsAiming)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(msg.AimingIndex))
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandAiming) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadBool(buffer, &msg.IsAiming)
	if err != nil {
		return err
	}
	var val int32
	err = helpers.ReadInt32(buffer, &val)
	msg.AimingIndex = int(val)
	if err != nil {
		return err
	}
	return nil
}
