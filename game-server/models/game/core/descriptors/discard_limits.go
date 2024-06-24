package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1501
type DiscardLimitsDescriptor struct {
	Items map[string]int
}

func (fastAccess *DiscardLimitsDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteInt32(buffer, int32(len(fastAccess.Items)))
	if err != nil {
		return err
	}
	for k, v := range fastAccess.Items {
		err = helpers.WriteUTF16String(buffer, k)
		if err != nil {
			return err
		}
		err = helpers.WriteInt32(buffer, int32(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func (fastAccess *DiscardLimitsDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var ItemsLength int32
	err = helpers.ReadInt32(buffer, &ItemsLength)
	if err != nil {
		return err
	}
	fastAccess.Items = make(map[string]int)
	for i := 0; i < int(ItemsLength); i++ {
		var k string
		err = helpers.ReadUTF16String(buffer, &k)
		if err != nil {
			return err
		}
		var v int32
		err = helpers.ReadInt32(buffer, &v)
		if err != nil {
			return err
		}
		fastAccess.Items[k] = int(v)
	}
	return nil
}
