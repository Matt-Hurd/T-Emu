package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type RpcSyncLighthouseTraderZoneData struct {
	data GStruct133
}

func (rsp *RpcSyncLighthouseTraderZoneData) Deserialize(buf *bytes.Buffer) error {
	if err := rsp.data.Deserialize(buf); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSyncLighthouseTraderZoneData) Serialize(buf *bytes.Buffer) error {
	if err := rsp.data.Serialize(buf); err != nil {
		return err
	}
	return nil
}

type GStruct383 struct {
	Nickname string
	Status   enums.RadioTransmitterStatus
}

func (g *GStruct383) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, g.Nickname); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(g.Status)); err != nil {
		return err
	}
	return nil
}

func (g *GStruct383) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &g.Nickname); err != nil {
		return err
	}
	var status int32
	if err := helpers.ReadInt32(buf, &status); err != nil {
		return err
	}
	return nil
}

type GStruct133 struct {
	AllowedPlayers   []GStruct383
	UnallowedPlayers []GStruct383
}

func (g *GStruct133) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt16(buf, int16(len(g.AllowedPlayers))); err != nil {
		return err
	}
	for _, v := range g.AllowedPlayers {
		if err := v.Serialize(buf); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt16(buf, int16(len(g.UnallowedPlayers))); err != nil {
		return err
	}
	for _, v := range g.UnallowedPlayers {
		if err := v.Serialize(buf); err != nil {
			return err
		}
	}
	return nil
}

func (g *GStruct133) Deserialize(buf *bytes.Buffer) error {
	var length int16
	if err := helpers.ReadInt16(buf, &length); err != nil {
		return err
	}
	for i := 0; i < int(length); i++ {
		var v GStruct383
		if err := v.Deserialize(buf); err != nil {
			return err
		}
		g.AllowedPlayers = append(g.AllowedPlayers, v)
	}
	if err := helpers.ReadInt16(buf, &length); err != nil {
		return err
	}
	for i := 0; i < int(length); i++ {
		var v GStruct383
		if err := v.Deserialize(buf); err != nil {
			return err
		}
		g.UnallowedPlayers = append(g.UnallowedPlayers, v)
	}
	return nil
}
