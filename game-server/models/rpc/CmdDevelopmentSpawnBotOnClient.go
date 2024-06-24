package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type CmdDevelopmentSpawnBotOnClient struct {
	side       enums.PlayerSide
	instanceId int32
}

func (rsp *CmdDevelopmentSpawnBotOnClient) Deserialize(buf *bytes.Buffer) error {
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.side = enums.PlayerSide(val)
	}
	if err := helpers.ReadInt32(buf, &rsp.instanceId); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdDevelopmentSpawnBotOnClient) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, int32(rsp.side)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.instanceId); err != nil {
		return err
	}
	return nil
}
