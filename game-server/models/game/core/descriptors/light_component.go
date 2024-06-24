package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type LightComponentDescriptor struct {
	IsActive     bool
	SelectedMode int32
}

func (lightComponent *LightComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, lightComponent.IsActive); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, lightComponent.SelectedMode); err != nil {
		return err
	}
	return nil
}

func (lightComponent *LightComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadBool(buffer, &lightComponent.IsActive); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buffer, &lightComponent.SelectedMode); err != nil {
		return err
	}
	return nil
}
