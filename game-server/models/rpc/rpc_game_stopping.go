package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameStopping struct {
	netId uint32
}

func (rsp *RpcGameStopping) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *RpcGameStopping) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
