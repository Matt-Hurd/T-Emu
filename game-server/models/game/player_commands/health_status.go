package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandHealthStatus struct {
	IsAlive      bool
	HealthStatus enums.TagStatus
}

func (msg *PlayerCommandHealthStatus) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteBool(buffer, msg.IsAlive)
	if err != nil {
		return err
	}
	err = helpers.WriteUInt32(buffer, uint32(msg.HealthStatus))
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandHealthStatus) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadBool(buffer, &msg.IsAlive)
	if err != nil {
		return err
	}
	var val uint32
	err = helpers.ReadUInt32(buffer, &val)
	msg.HealthStatus = enums.TagStatus(val)
	if err != nil {
		return err
	}
	return nil
}
