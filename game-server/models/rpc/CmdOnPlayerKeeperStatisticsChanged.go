package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type CmdOnPlayerKeeperStatisticsChanged struct {
	playerProfileID string
	statisticsType  enums.CounterTag
	valueToSet      int32
}

func (rsp *CmdOnPlayerKeeperStatisticsChanged) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.playerProfileID); err != nil {
		return err
	}
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		rsp.statisticsType = enums.CounterTag(val)
	}
	if err := helpers.ReadInt32(buf, &rsp.valueToSet); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdOnPlayerKeeperStatisticsChanged) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.playerProfileID); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(rsp.statisticsType)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, rsp.valueToSet); err != nil {
		return err
	}
	return nil
}
