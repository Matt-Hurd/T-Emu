package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type TagComponentDescriptor struct {
	Name  string
	Color int32
}

func (tagComponent *TagComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, tagComponent.Name)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, tagComponent.Color)
	if err != nil {
		return err
	}
	return nil
}

func (tagComponent *TagComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &tagComponent.Name)
	if err != nil {
		return err
	}
	err = helpers.ReadInt32(buffer, &tagComponent.Color)
	if err != nil {
		return err
	}
	return nil
}
