package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type CmdDevelopmentSpawnBotOnServer struct {
	side       enums.EPlayerSide
	profile    enums.WildSpawnType
	difficulty enums.BotDifficulty
}

func (rsp *CmdDevelopmentSpawnBotOnServer) Deserialize(buf *bytes.Buffer) error {
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.side = enums.EPlayerSide(val)
	}
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.profile = enums.WildSpawnType(val)
	}
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.difficulty = enums.BotDifficulty(val)
	}
	return nil
}

func (rsp *CmdDevelopmentSpawnBotOnServer) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, int32(rsp.side)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(rsp.profile)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(rsp.difficulty)); err != nil {
		return err
	}
	return nil
}
