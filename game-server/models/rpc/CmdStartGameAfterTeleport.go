package rpc

import (
	"bytes"
)

type CmdStartGameAfterTeleport struct {
}

func (rsp *CmdStartGameAfterTeleport) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdStartGameAfterTeleport) Serialize(buf *bytes.Buffer) error {
	return nil
}
