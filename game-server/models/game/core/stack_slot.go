package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1508
type StackSlot struct {
	Id             string
	ContainedItems []Item
}

func (stackSlot *StackSlot) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, stackSlot.Id)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(len(stackSlot.ContainedItems)))
	if err != nil {
		return err
	}
	for _, v := range stackSlot.ContainedItems {
		err = v.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (stackSlot *StackSlot) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &stackSlot.Id)
	if err != nil {
		return err
	}
	var ContainedItemsLength int32
	err = helpers.ReadInt32(buffer, &ContainedItemsLength)
	if err != nil {
		return err
	}
	stackSlot.ContainedItems = make([]Item, ContainedItemsLength)
	for i := range stackSlot.ContainedItems {
		err = stackSlot.ContainedItems[i].Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
