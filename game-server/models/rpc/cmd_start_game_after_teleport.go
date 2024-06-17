package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdStartGameAfterTeleport struct {
	netId uint32
}

func (rsp *CmdStartGameAfterTeleport) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *CmdStartGameAfterTeleport) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
