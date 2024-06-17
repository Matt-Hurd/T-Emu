package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdStartGame struct {
	netId uint32
}

func (rsp *CmdStartGame) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *CmdStartGame) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
