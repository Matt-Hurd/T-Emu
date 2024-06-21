package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1518
type MedKitComponent struct {
	HpPercent float32
}

func (component *MedKitComponent) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, component.HpPercent)
	if err != nil {
		return err
	}
	return nil
}

func (component *MedKitComponent) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &component.HpPercent)
	if err != nil {
		return err
	}
	return nil
}
