package rpc

import (
	"bytes"
)

type CmdStopGame struct {
}

func (rsp *CmdStopGame) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdStopGame) Serialize(buf *bytes.Buffer) error {
	return nil
}
