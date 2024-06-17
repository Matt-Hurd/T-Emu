package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdGetTraderServicesData struct {
	playerProfileID string
	traderId        string
}

func (rsp *CmdGetTraderServicesData) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.playerProfileID); err != nil {
		return err
	}
	if err := helpers.ReadString(buf, &rsp.traderId); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdGetTraderServicesData) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.playerProfileID); err != nil {
		return err
	}
	if err := helpers.WriteString(buf, rsp.traderId); err != nil {
		return err
	}
	return nil
}
