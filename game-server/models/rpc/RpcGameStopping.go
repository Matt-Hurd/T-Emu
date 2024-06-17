package rpc

import (
	"bytes"
)

type RpcGameStopping struct {
}

func (rsp *RpcGameStopping) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *RpcGameStopping) Serialize(buf *bytes.Buffer) error {
	return nil
}
