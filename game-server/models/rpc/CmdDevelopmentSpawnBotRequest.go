package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type CmdDevelopmentSpawnBotRequest struct {
	side enums.EPlayerSide
}

func (rsp *CmdDevelopmentSpawnBotRequest) Deserialize(buf *bytes.Buffer) error {
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.side = enums.EPlayerSide(val)
	}
	return nil
}

func (rsp *CmdDevelopmentSpawnBotRequest) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, int32(rsp.side)); err != nil {
		return err
	}
	return nil
}
