package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdPlayerEffectsPause struct {
	playerProfileID string
	isPaused        bool
}

func (rsp *CmdPlayerEffectsPause) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.playerProfileID); err != nil {
		return err
	}
	if err := helpers.ReadBool(buf, &rsp.isPaused); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdPlayerEffectsPause) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.playerProfileID); err != nil {
		return err
	}
	if err := helpers.WriteBool(buf, rsp.isPaused); err != nil {
		return err
	}
	return nil
}
