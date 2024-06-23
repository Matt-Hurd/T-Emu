package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetUnderRoofStatus struct {
	hasValue    bool
	IsUnderRoof bool
}

func (msg *PlayerCommandSetUnderRoofStatus) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.hasValue); err != nil {
		return err
	}
	if msg.hasValue {
		if err := helpers.WriteBool(buffer, msg.IsUnderRoof); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandSetUnderRoofStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.hasValue); err != nil {
		return err
	}
	if msg.hasValue {
		if err = helpers.ReadBool(buffer, &msg.IsUnderRoof); err != nil {
			return err
		}
	}
	return nil
}
