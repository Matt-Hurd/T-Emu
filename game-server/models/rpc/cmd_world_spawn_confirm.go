package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdWorldSpawnConfirm struct {
	netId uint32
}

func (rsp *CmdWorldSpawnConfirm) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *CmdWorldSpawnConfirm) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
