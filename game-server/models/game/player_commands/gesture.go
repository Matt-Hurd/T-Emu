package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandGesture struct {
	Gesture enums.Gesture
}

func (msg *PlayerCommandGesture) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.Gesture)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandGesture) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.Gesture)); err != nil {
		return err
	}
	return nil
}
