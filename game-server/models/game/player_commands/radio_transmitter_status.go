package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandRadioTransmitterStatus struct {
	RadioTransmitterStatus enums.RadioTransmitterStatus
}

func (msg *PlayerCommandRadioTransmitterStatus) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.RadioTransmitterStatus)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandRadioTransmitterStatus) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.RadioTransmitterStatus)); err != nil {
		return err
	}
	return nil
}
