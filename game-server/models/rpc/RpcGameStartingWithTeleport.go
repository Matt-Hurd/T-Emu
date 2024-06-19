package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type RpcGameStartingWithTeleport struct {
	position       core.Vector3
	exfiltrationId int32
	entryPoint     string
}

func (rsp *RpcGameStartingWithTeleport) Deserialize(buf *bytes.Buffer) error {
	if err := rsp.position.Deserialize(buf); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buf, &rsp.exfiltrationId); err != nil {
		return err
	}
	if err := helpers.ReadString(buf, &rsp.entryPoint); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcGameStartingWithTeleport) Serialize(buf *bytes.Buffer) error {
	if err := rsp.position.Serialize(buf); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.exfiltrationId); err != nil {
		return err
	}
	if err := helpers.WriteString(buf, rsp.entryPoint); err != nil {
		return err
	}
	return nil
}
