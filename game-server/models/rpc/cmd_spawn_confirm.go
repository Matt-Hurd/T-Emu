package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdSpawnConfirm struct {
	netId    uint32
	playerId int32
}

func (rsp *CmdSpawnConfirm) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadUInt32(buf, &rsp.netId); err != nil {
		return err
	}
	return helpers.ReadInt32(buf, &rsp.playerId)
}

func (rsp *CmdSpawnConfirm) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteUInt32(buf, rsp.netId); err != nil {
		return err
	}
	return helpers.WriteInt32(buf, rsp.playerId)
}
