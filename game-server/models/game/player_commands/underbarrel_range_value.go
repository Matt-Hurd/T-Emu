package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandUnderbarrelRangeValue struct {
	RangeValue     int32
	LastRangeValue int32
}

func (msg *PlayerCommandUnderbarrelRangeValue) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.RangeValue); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, msg.LastRangeValue); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandUnderbarrelRangeValue) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.RangeValue); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &msg.LastRangeValue); err != nil {
		return err
	}
	return nil
}
