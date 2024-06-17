package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdSpawnConfirm struct {
	playerId int32
}

func (rsp *CmdSpawnConfirm) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt32(buf, &rsp.playerId); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdSpawnConfirm) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, rsp.playerId); err != nil {
		return err
	}
	return nil
}
