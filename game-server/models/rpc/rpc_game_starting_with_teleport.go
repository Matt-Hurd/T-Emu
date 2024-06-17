package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type RpcGameStartingWithTeleport struct {
	netId      uint32
	Position   core.Vector3
	exfilId    int32
	EntryPoint string
}

func (rsp *RpcGameStartingWithTeleport) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadUInt32(buf, &rsp.netId); err != nil {
		return err
	}
	if vec, err := core.DeserializeVector3(buf); err != nil {
		return err
	} else {
		rsp.Position = vec
	}
	if err := helpers.ReadInt32(buf, &rsp.exfilId); err != nil {
		return err
	}
	return helpers.ReadString(buf, &rsp.EntryPoint)
}

func (rsp *RpcGameStartingWithTeleport) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteUInt32(buf, rsp.netId); err != nil {
		return err
	}
	if err := rsp.Position.Serialize(buf); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.exfilId); err != nil {
		return err
	}
	return helpers.WriteString(buf, rsp.EntryPoint)
}
