package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandSetUnderRoofStatus struct {
	IsUnderRoof *bool
}

func (msg *PlayerCommandSetUnderRoofStatus) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.IsUnderRoof == nil); err != nil {
		return err
	}
	if msg.IsUnderRoof != nil {
		if err := helpers.WriteBool(buffer, *msg.IsUnderRoof); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandSetUnderRoofStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var isNil bool
	if err = helpers.ReadBool(buffer, &isNil); err != nil {
		return err
	}
	if !isNil {
		if err = helpers.ReadBool(buffer, msg.IsUnderRoof); err != nil {
			return err
		}
	}
	return nil
}
