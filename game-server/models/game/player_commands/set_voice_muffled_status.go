package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetVoiceMuffledStatus struct {
	hasValue       bool
	IsVoiceMuffled bool
}

func (msg *PlayerCommandSetVoiceMuffledStatus) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.hasValue); err != nil {
		return err
	}
	if msg.hasValue {
		if err := helpers.WriteBool(buffer, msg.IsVoiceMuffled); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandSetVoiceMuffledStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.hasValue); err != nil {
		return err
	}
	if msg.hasValue {
		if err = helpers.ReadBool(buffer, &msg.IsVoiceMuffled); err != nil {
			return err
		}
	}
	return nil
}
