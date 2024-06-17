package rpc

import (
	"bytes"
)

type CmdRestartGameInitiate struct {
}

func (rsp *CmdRestartGameInitiate) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdRestartGameInitiate) Serialize(buf *bytes.Buffer) error {
	return nil
}
