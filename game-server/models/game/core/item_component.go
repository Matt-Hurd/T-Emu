package core

import (
	"bytes"
	"game-server/helpers"
)

type ItemComponent struct {
	Id           string `json:"_id"`
	Tpl          string `json:"_tpl"`
	ParentId     string
	SlotId       string
	UpdJson      string // Technically this should be deserialized
	LocationJson string
}

func (itemComponent *ItemComponent) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteMongoId(buffer, itemComponent.Id); err != nil {
		return err
	}
	if err := helpers.WriteMongoId(buffer, itemComponent.Tpl); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, len(itemComponent.ParentId) > 0); err != nil {
		return err
	}
	if len(itemComponent.ParentId) > 0 {
		if err := helpers.WriteMongoId(buffer, itemComponent.ParentId); err != nil {
			return err
		}
	}
	if err := helpers.WriteString(buffer, itemComponent.SlotId); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, len(itemComponent.UpdJson) > 0); err != nil {
		return err
	}
	if len(itemComponent.UpdJson) > 0 {
		if err := helpers.WriteString(buffer, itemComponent.UpdJson); err != nil {
			return err
		}
	}
	if err := helpers.WriteBool(buffer, len(itemComponent.LocationJson) > 0); err != nil {
		return err
	}
	if len(itemComponent.LocationJson) > 0 {
		if err := helpers.WriteString(buffer, itemComponent.LocationJson); err != nil {
			return err
		}
	}
	return nil
}

func (itemComponent *ItemComponent) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadMongoId(buffer, &itemComponent.Id); err != nil {
		return err
	}
	if err = helpers.ReadMongoId(buffer, &itemComponent.Tpl); err != nil {
		return err
	}
	var hasParentId bool
	if err = helpers.ReadBool(buffer, &hasParentId); err != nil {
		return err
	}
	if hasParentId {
		if err = helpers.ReadMongoId(buffer, &itemComponent.ParentId); err != nil {
			return err
		}
	}
	if err = helpers.ReadString(buffer, &itemComponent.SlotId); err != nil {
		return err
	}
	var hasUpdJson bool
	if err = helpers.ReadBool(buffer, &hasUpdJson); err != nil {
		return err
	}
	if hasUpdJson {
		if err = helpers.ReadString(buffer, &itemComponent.UpdJson); err != nil {
			return err
		}
	}
	var hasLocationJson bool
	if err = helpers.ReadBool(buffer, &hasLocationJson); err != nil {
		return err
	}
	if hasLocationJson {
		if err = helpers.ReadString(buffer, &itemComponent.LocationJson); err != nil {
			return err
		}
	}
	return nil
}
