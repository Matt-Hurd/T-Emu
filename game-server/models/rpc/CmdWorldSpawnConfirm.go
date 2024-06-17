package rpc

import (
	"bytes"
)

type CmdWorldSpawnConfirm struct {
}

func (rsp *CmdWorldSpawnConfirm) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdWorldSpawnConfirm) Serialize(buf *bytes.Buffer) error {
	return nil
}
