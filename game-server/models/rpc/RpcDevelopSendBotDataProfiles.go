package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcDevelopSendBotDataProfiles struct {
	data []byte
}

func (rsp *RpcDevelopSendBotDataProfiles) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcDevelopSendBotDataProfiles) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
