package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandMedEffectStatus struct {
	EffectAdded bool
	EffectId    string
}

func (msg *PlayerCommandMedEffectStatus) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteBool(buffer, msg.EffectAdded)
	if err != nil {
		return err
	}
	err = helpers.WriteString(buffer, msg.EffectId)
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandMedEffectStatus) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadBool(buffer, &msg.EffectAdded)
	if err != nil {
		return err
	}
	err = helpers.ReadString(buffer, &msg.EffectId)
	if err != nil {
		return err
	}
	return nil
}
