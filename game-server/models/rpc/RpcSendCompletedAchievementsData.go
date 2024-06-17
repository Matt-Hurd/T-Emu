package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type RpcSendCompletedAchievementsData struct {
	data GStruct220
}

func (rsp *RpcSendCompletedAchievementsData) Deserialize(buf *bytes.Buffer) error {
	if err := rsp.data.Deserialize(buf); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSendCompletedAchievementsData) Serialize(buf *bytes.Buffer) error {
	if err := rsp.data.Serialize(buf); err != nil {
		return err
	}
	return nil
}

type GStruct221 struct {
	QuestId     int32
	Status      enums.EQuestStatus
	ConditionId int32
	Value       float64
	Notify      bool
}

func (g *GStruct221) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, g.QuestId); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(g.Status)); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, g.ConditionId); err != nil {
		return err
	}
	if err := helpers.WriteFloat64(buf, g.Value); err != nil {
		return err
	}
	if err := helpers.WriteBool(buf, g.Notify); err != nil {
		return err
	}
	return nil
}

func (g *GStruct221) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt32(buf, &g.QuestId); err != nil {
		return err
	}
	var status int32
	if err := helpers.ReadInt32(buf, &status); err != nil {
		return err
	}
	g.Status = enums.EQuestStatus(status)
	if err := helpers.ReadInt32(buf, &g.ConditionId); err != nil {
		return err
	}
	if err := helpers.ReadFloat64(buf, &g.Value); err != nil {
		return err
	}
	if err := helpers.ReadBool(buf, &g.Notify); err != nil {
		return err
	}
	return nil
}

type GStruct220 struct {
	Data []GStruct221
}

func (g *GStruct220) Serialize(buf *bytes.Buffer) error {
	length := int16(len(g.Data))
	if err := helpers.WriteInt16(buf, length); err != nil {
		return err
	}
	for _, item := range g.Data {
		if err := item.Serialize(buf); err != nil {
			return err
		}
	}
	return nil
}

func (g *GStruct220) Deserialize(buf *bytes.Buffer) error {
	var length int16
	if err := helpers.ReadInt16(buf, &length); err != nil {
		return err
	}
	g.Data = make([]GStruct221, length)
	for i := 0; i < int(length); i++ {
		var item GStruct221
		if err := item.Deserialize(buf); err != nil {
			return err
		}
		g.Data[i] = item
	}
	return nil
}
