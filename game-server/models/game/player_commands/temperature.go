package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandTemperature struct {
	Temperature float32
}

func (msg *PlayerCommandTemperature) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, msg.Temperature)
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandTemperature) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &msg.Temperature)
	if err != nil {
		return err
	}
	return nil
}
