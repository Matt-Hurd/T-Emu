package rpc

import (
	"bytes"
)

type CmdStartGame struct {
}

func (rsp *CmdStartGame) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdStartGame) Serialize(buf *bytes.Buffer) error {
	return nil
}
