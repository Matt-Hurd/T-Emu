package rpc

import (
	"bytes"
)

type CmdRestartGame struct {
}

func (rsp *CmdRestartGame) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdRestartGame) Serialize(buf *bytes.Buffer) error {
	return nil
}
