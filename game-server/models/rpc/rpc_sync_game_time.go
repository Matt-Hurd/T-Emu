package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSyncGameTime struct {
	netId uint32
	Time  int64
}

func (rsp *RpcSyncGameTime) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadUInt32(buf, &rsp.netId); err != nil {
		return err
	}
	return helpers.ReadInt64(buf, &rsp.Time)
}

func (rsp *RpcSyncGameTime) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteUInt32(buf, rsp.netId); err != nil {
		return err
	}
	return helpers.WriteInt64(buf, rsp.Time)
}
