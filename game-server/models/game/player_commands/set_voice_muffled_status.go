package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetVoiceMuffledStatus struct {
	IsVoiceMuffled *bool
}

func (msg *PlayerCommandSetVoiceMuffledStatus) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.IsVoiceMuffled == nil); err != nil {
		return err
	}
	if msg.IsVoiceMuffled != nil {
		if err := helpers.WriteBool(buffer, *msg.IsVoiceMuffled); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandSetVoiceMuffledStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var isNil bool
	if err = helpers.ReadBool(buffer, &isNil); err != nil {
		return err
	}
	if !isNil {
		if err = helpers.ReadBool(buffer, msg.IsVoiceMuffled); err != nil {
			return err
		}
	}
	return nil
}
