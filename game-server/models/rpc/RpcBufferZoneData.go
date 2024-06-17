package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcBufferZoneData struct {
	data []byte
}

func (rsp *RpcBufferZoneData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcBufferZoneData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
