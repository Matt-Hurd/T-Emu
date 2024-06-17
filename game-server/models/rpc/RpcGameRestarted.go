package rpc

import (
	"bytes"
)

type RpcGameRestarted struct {
}

func (rsp *RpcGameRestarted) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *RpcGameRestarted) Serialize(buf *bytes.Buffer) error {
	return nil
}
