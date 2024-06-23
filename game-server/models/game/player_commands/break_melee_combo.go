package playercommands

import (
	"bytes"
)

type PlayerCommandBreakMeleeCombo struct {
}

func (msg *PlayerCommandBreakMeleeCombo) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandBreakMeleeCombo) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
