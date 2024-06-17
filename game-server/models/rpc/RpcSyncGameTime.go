package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSyncGameTime struct {
	time int64
}

func (rsp *RpcSyncGameTime) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt64(buf, &rsp.time); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSyncGameTime) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt64(buf, rsp.time); err != nil {
		return err
	}
	return nil
}
