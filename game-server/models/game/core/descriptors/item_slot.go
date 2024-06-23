package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1503
type SlotDescriptor struct {
	Id            string
	ContainedItem ItemDescriptor
}

func (itemSlot *SlotDescriptor) Serialize(buffer *bytes.Buffer) error {
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

func (itemSlot *SlotDescriptor) Deserialize(buffer *bytes.Buffer) error {
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
