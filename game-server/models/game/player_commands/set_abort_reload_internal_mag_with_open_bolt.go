package playercommands

import (
	"bytes"
)

type PlayerCommandSetAbortReloadInternalMagWithOpenBolt struct {
}

func (msg *PlayerCommandSetAbortReloadInternalMagWithOpenBolt) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandSetAbortReloadInternalMagWithOpenBolt) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
