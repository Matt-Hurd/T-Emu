package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcDevelopSendBotDataGroups struct {
	data []byte
}

func (rsp *RpcDevelopSendBotDataGroups) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcDevelopSendBotDataGroups) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
