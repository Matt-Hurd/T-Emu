package rpc

import (
	"bytes"
)

type RpcGameSpawned struct {
}

func (rsp *RpcGameSpawned) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *RpcGameSpawned) Serialize(buf *bytes.Buffer) error {
	return nil
}
