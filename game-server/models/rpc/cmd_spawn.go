package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdSpawn struct {
	netId uint32
}

func (rsp *CmdSpawn) Deserialize(buf *bytes.Buffer) error {
	return helpers.ReadUInt32(buf, &rsp.netId)
}

func (rsp *CmdSpawn) Serialize(buf *bytes.Buffer) error {
	return helpers.WriteUInt32(buf, rsp.netId)
}
