package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameSpawned struct {
	netId uint32
}

func DeserializeRpcGameSpawned(buf *bytes.Buffer) (*RpcGameSpawned, error) {
	rsp := &RpcGameSpawned{}
	err := helpers.ReadUInt32(buf, &rsp.netId)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (rsp *RpcGameSpawned) Serialize(buf *bytes.Buffer) error {
	err := helpers.WriteUInt32(buf, rsp.netId)
	if err != nil {
		return err
	}
	return nil
}
