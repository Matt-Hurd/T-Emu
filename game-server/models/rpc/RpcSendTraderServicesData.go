package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSendTraderServicesData struct {
	packet []byte
}

func (rsp *RpcSendTraderServicesData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.packet); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSendTraderServicesData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.packet); err != nil {
		return err
	}
	return nil
}
