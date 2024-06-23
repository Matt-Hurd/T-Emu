package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1518
type ResourceItemComponentDescriptor struct {
	Resource float32
}

func (component *ResourceItemComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, component.Resource)
	if err != nil {
		return err
	}
	return nil
}

func (component *ResourceItemComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &component.Resource)
	if err != nil {
		return err
	}
	return nil
}
