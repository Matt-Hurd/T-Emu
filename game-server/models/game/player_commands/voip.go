package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandVoIP struct {
	VoIPState enums.VoipState
}

func (msg *PlayerCommandVoIP) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, byte(msg.VoIPState)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandVoIP) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var VoIPState byte
	if err = helpers.ReadByte(buffer, &VoIPState); err != nil {
		return err
	}
	msg.VoIPState = enums.VoipState(VoIPState)
	return nil
}
