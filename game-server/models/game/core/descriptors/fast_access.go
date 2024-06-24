package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1501
type FastAccessDescriptor struct {
	Items map[int]string
}

func (fastAccess *FastAccessDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteInt32(buffer, int32(len(fastAccess.Items)))
	if err != nil {
		return err
	}
	for k, v := range fastAccess.Items {
		err = helpers.WriteInt32(buffer, int32(k))
		if err != nil {
			return err
		}
		err = helpers.WriteUTF16String(buffer, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fastAccess *FastAccessDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var ItemsLength int32
	err = helpers.ReadInt32(buffer, &ItemsLength)
	if err != nil {
		return err
	}
	fastAccess.Items = make(map[int]string)
	for i := 0; i < int(ItemsLength); i++ {
		var k int32
		err = helpers.ReadInt32(buffer, &k)
		if err != nil {
			return err
		}
		var v string
		err = helpers.ReadUTF16String(buffer, &v)
		if err != nil {
			return err
		}
		fastAccess.Items[int(k)] = v
	}
	return nil
}
