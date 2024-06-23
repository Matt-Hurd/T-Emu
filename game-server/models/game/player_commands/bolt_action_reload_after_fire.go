package playercommands

import "bytes"

type PlayerCommandBoltActionReloadAfterFire struct {
}

func (msg *PlayerCommandBoltActionReloadAfterFire) Serialize(buffer *bytes.Buffer) error {
	return nil
}

func (msg *PlayerCommandBoltActionReloadAfterFire) Deserialize(buffer *bytes.Buffer) error {
	return nil
}
