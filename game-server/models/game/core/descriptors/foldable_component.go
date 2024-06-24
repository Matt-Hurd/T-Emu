package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type FoldableComponentDescriptor struct {
	Folded bool
}

func (msg *FoldableComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.Folded); err != nil {
		return err
	}
	return nil
}

func (msg *FoldableComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadBool(buffer, &msg.Folded); err != nil {
		return err
	}
	return nil
}
