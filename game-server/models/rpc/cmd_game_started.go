package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdGameStarted struct {
	netId uint32
}

func (rsp *CmdGameStarted) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *CmdGameStarted) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
