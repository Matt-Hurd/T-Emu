package rpc

import (
	"bytes"
)

type RpcGameRestarting struct {
}

func (rsp *RpcGameRestarting) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *RpcGameRestarting) Serialize(buf *bytes.Buffer) error {
	return nil
}
