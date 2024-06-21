package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1503
type ItemSlot struct {
	Id            string
	ContainedItem Item
}

func (itemSlot *ItemSlot) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, itemSlot.Id)
	if err != nil {
		return err
	}
	err = itemSlot.ContainedItem.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (itemSlot *ItemSlot) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &itemSlot.Id)
	if err != nil {
		return err
	}
	err = itemSlot.ContainedItem.Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}
