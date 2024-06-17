package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcSoftStopNotification struct {
	sessionSeconds int32
}

func (rsp *RpcSoftStopNotification) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt32(buf, &rsp.sessionSeconds); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSoftStopNotification) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, rsp.sessionSeconds); err != nil {
		return err
	}
	return nil
}
