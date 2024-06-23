package core

import (
	"bytes"
	"game-server/helpers"
)

type ComponentialItem struct {
	Id         string
	Components []ItemComponent
}

func (componentialItem *ComponentialItem) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, len(componentialItem.Id) > 0); err != nil {
		return err
	} else if len(componentialItem.Id) == 0 {
		return nil
	}
	if err := helpers.WriteMongoId(buffer, componentialItem.Id); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(len(componentialItem.Components))); err != nil {
		return err
	}
	for _, component := range componentialItem.Components {
		if err := component.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (componentialItem *ComponentialItem) Deserialize(buffer *bytes.Buffer) error {
	var hasId bool
	if err := helpers.ReadBool(buffer, &hasId); err != nil {
		return err
	} else if !hasId {
		return nil
	}
	if err := helpers.ReadMongoId(buffer, &componentialItem.Id); err != nil {
		return err
	}
	var componentsCount byte
	if err := helpers.ReadByte(buffer, &componentsCount); err != nil {
		return err
	}
	componentialItem.Components = make([]ItemComponent, componentsCount)
	for i := 0; i < int(componentsCount); i++ {
		if err := componentialItem.Components[i].Deserialize(buffer); err != nil {
			return err
		}
	}
	return nil
}
