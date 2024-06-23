package playercommands

import "bytes"

type PlayerCommandDryShotFirearm struct {
}

func (msg *PlayerCommandDryShotFirearm) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandDryShotFirearm) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
