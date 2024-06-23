package playercommands

import (
	"bytes"
)

type PlayerCommandStopSearchContent struct {
}

func (msg *PlayerCommandStopSearchContent) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandStopSearchContent) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
