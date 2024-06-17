package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameStarted struct {
	pastTime       float32
	sessionSeconds int32
}

func (rsp *RpcGameStarted) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadFloat32(buf, &rsp.pastTime); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buf, &rsp.sessionSeconds); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcGameStarted) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteFloat32(buf, rsp.pastTime); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.sessionSeconds); err != nil {
		return err
	}
	return nil
}
