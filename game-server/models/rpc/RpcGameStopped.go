package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type RpcGameStopped struct {
	exitStatus enums.ExitStatus
	playTime   int32
}

func (rsp *RpcGameStopped) Deserialize(buf *bytes.Buffer) error {
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.exitStatus = enums.ExitStatus(val)
	}
	if err := helpers.ReadInt32(buf, &rsp.playTime); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcGameStopped) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, int32(rsp.exitStatus)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.playTime); err != nil {
		return err
	}
	return nil
}
