package rpc

import (
	"bytes"
)

type CmdGameStarted struct {
}

func (rsp *CmdGameStarted) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdGameStarted) Serialize(buf *bytes.Buffer) error {
	return nil
}
