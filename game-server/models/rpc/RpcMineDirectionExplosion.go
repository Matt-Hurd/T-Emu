package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcMineDirectionExplosion struct {
	data []byte
}

func (rsp *RpcMineDirectionExplosion) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcMineDirectionExplosion) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
