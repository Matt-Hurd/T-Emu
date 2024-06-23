package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandMedEffectResource struct {
	MedItemID string
	Resource  float32
}

func (msg *PlayerCommandMedEffectResource) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.MedItemID); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.Resource); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandMedEffectResource) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.MedItemID); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.Resource); err != nil {
		return err
	}
	return nil
}
