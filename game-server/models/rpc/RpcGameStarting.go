package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameStarting struct {
	seconds int32
}

func (rsp *RpcGameStarting) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt32(buf, &rsp.seconds); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcGameStarting) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, rsp.seconds); err != nil {
		return err
	}
	return nil
}
