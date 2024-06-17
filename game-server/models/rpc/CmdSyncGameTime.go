package rpc

import (
	"bytes"
)

type CmdSyncGameTime struct {
}

func (rsp *CmdSyncGameTime) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdSyncGameTime) Serialize(buf *bytes.Buffer) error {
	return nil
}
