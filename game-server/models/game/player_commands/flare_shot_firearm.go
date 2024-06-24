package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/math"
)

type PlayerCommandFlareShotFirearm struct {
	ShotPosition  math.Vector3
	ShotDirection math.Vector3
	AmmoTemplate  string
}

func (msg *PlayerCommandFlareShotFirearm) Serialize(buffer *bytes.Buffer) error {
	if err := msg.ShotPosition.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.ShotDirection.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandFlareShotFirearm) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = msg.ShotPosition.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.ShotDirection.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &msg.AmmoTemplate); err != nil {
		return err
	}
	return nil
}
