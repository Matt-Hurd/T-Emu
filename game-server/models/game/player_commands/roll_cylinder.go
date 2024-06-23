package playercommands

import (
	"bytes"
)

type PlayerCommandRollCylinder struct {
}

func (msg *PlayerCommandRollCylinder) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandRollCylinder) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
