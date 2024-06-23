package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetAnimationLayerWeight struct {
	LayerIndex int32
	Weight     float32
}

func (msg *PlayerCommandSetAnimationLayerWeight) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.LayerIndex); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.Weight); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSetAnimationLayerWeight) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.LayerIndex); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.Weight); err != nil {
		return err
	}
	return nil
}
