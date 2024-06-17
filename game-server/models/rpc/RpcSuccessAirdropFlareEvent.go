package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSuccessAirdropFlareEvent struct {
	canSendAirdrop bool
}

func (rsp *RpcSuccessAirdropFlareEvent) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadBool(buf, &rsp.canSendAirdrop); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSuccessAirdropFlareEvent) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteBool(buf, rsp.canSendAirdrop); err != nil {
		return err
	}
	return nil
}
