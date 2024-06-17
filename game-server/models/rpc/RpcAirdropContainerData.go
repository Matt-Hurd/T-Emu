package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcAirdropContainerData struct {
	data []byte
}

func (rsp *RpcAirdropContainerData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBytesAndSize(buf, &rsp.data); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcAirdropContainerData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBytesAndSize(buf, rsp.data); err != nil {
		return err
	}
	return nil
}
