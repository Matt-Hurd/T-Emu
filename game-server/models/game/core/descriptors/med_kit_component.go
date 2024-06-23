package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1518
type MedKitComponentDescriptor struct {
	HpPercent float32
}

func (component *MedKitComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, component.HpPercent)
	if err != nil {
		return err
	}
	return nil
}

func (component *MedKitComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &component.HpPercent)
	if err != nil {
		return err
	}
	return nil
}
