package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcDevelopSendBotDataZone struct {
	data []byte
}

func (rsp *RpcDevelopSendBotDataZone) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcDevelopSendBotDataZone) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
