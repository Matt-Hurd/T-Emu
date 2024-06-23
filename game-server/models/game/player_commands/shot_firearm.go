package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PlayerCommandShotFirearm struct {
	ShotPosition  core.Vector3
	ShotDirection core.Vector3
	ChamberIndex  int32
	AmmoTemplate  string
}

func (msg *PlayerCommandShotFirearm) Serialize(buffer *bytes.Buffer) error {
	if err := msg.ShotPosition.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.ShotDirection.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.ChamberIndex); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandShotFirearm) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = msg.ShotPosition.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.ShotDirection.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.ChamberIndex); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}
