package rpc

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type RpcSendClientRadioTransmitterData struct {
	data GStruct134
}

func (rsp *RpcSendClientRadioTransmitterData) Deserialize(buf *bytes.Buffer) error {
	if err := rsp.data.Deserialize(buf); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSendClientRadioTransmitterData) Serialize(buf *bytes.Buffer) error {
	if err := rsp.data.Serialize(buf); err != nil {
		return err
	}
	return nil
}

type GStruct134 struct {
	PlayerProfileIDForSend string
	PlayerID               int32
	IsEncoded              bool
	Status                 enums.RadioTransmitterStatus
	IsAgressor             bool
}

// Serialize serializes the GStruct134 into a byte buffer.
func (g *GStruct134) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, g.PlayerProfileIDForSend); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, g.PlayerID); err != nil {
		return err
	}
	if err := helpers.WriteBool(buf, g.IsEncoded); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buf, int32(g.Status)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buf, g.IsAgressor); err != nil {
		return err
	}
	return nil
}

// Deserialize deserializes the byte buffer into GStruct134.
func (g *GStruct134) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &g.PlayerProfileIDForSend); err != nil {
		return err
	}
	if err := helpers.ReadInt32(buf, &g.PlayerID); err != nil {
		return err
	}
	if err := helpers.ReadBool(buf, &g.IsEncoded); err != nil {
		return err
	}
	var val int32
	if err := helpers.ReadInt32(buf, &val); err != nil {
		return err
	} else {
		g.Status = enums.RadioTransmitterStatus(val)
	}
	if err := helpers.ReadBool(buf, &g.IsAgressor); err != nil {
		return err
	}
	return nil
}
