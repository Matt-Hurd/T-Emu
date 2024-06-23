package playercommands

import (
	"bytes"
)

type PlayerCommandAbortReloadCylinderMagazine struct {
}

func (msg *PlayerCommandAbortReloadCylinderMagazine) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandAbortReloadCylinderMagazine) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
