package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcDevelopSendBotData struct {
	data []byte
}

func (rsp *RpcDevelopSendBotData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcDevelopSendBotData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
