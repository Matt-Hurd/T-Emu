package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdGetRadiotransmitterData struct {
	playerProfileID string
}

func (rsp *CmdGetRadiotransmitterData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.playerProfileID); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdGetRadiotransmitterData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.playerProfileID); err != nil {
		return err
	}
	return nil
}
