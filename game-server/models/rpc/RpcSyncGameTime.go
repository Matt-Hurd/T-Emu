package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSyncGameTime struct {
	Time uint64
}

func (rsp *RpcSyncGameTime) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadPackedUInt64(buf, &rsp.Time); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSyncGameTime) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WritePackedUInt64(buf, rsp.Time); err != nil {
		return err
	}
	return nil
}
