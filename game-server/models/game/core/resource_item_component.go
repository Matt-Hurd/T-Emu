package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1518
type ResourceItemComponent struct {
	Resource float32
}

func (component *ResourceItemComponent) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, component.Resource)
	if err != nil {
		return err
	}
	return nil
}

func (component *ResourceItemComponent) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &component.Resource)
	if err != nil {
		return err
	}
	return nil
}
