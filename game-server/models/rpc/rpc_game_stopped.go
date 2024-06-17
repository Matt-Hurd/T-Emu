package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameStopped struct {
	netId      uint32
	ExitStatus int32
	playTime   int32
}

func (rsp *RpcGameStopped) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadUInt32(buf, &rsp.netId); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buf, &rsp.ExitStatus); err != nil {
		return err
	}
	return helpers.ReadInt32(buf, &rsp.playTime)
}

func (rsp *RpcGameStopped) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteUInt32(buf, rsp.netId); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.ExitStatus); err != nil {
		return err
	}
	return helpers.WriteInt32(buf, rsp.playTime)
}
