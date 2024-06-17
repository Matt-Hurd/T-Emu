package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcGameMatching struct {
	activitiesCounter int16
	minCounter        int16
	seconds           int32
}

func (rsp *RpcGameMatching) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt16(buf, &rsp.activitiesCounter); err != nil {
		return err
	}
	if err := helpers.ReadInt16(buf, &rsp.minCounter); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buf, &rsp.seconds); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcGameMatching) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt16(buf, rsp.activitiesCounter); err != nil {
		return err
	}
	if err := helpers.WriteInt16(buf, rsp.minCounter); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.seconds); err != nil {
		return err
	}
	return nil
}
