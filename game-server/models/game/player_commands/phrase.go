package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandPhrase struct {
	PhraseId      int32
	PhraseCommand enums.PhraseTrigger
}

func (msg *PlayerCommandPhrase) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, msg.PhraseId); err != nil {
		return err
	}
	if err := buffer.WriteByte(byte(msg.PhraseCommand)); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandPhrase) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &msg.PhraseId); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&msg.PhraseCommand)); err != nil {
		return err
	}
	return nil
}
