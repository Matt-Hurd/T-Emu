package rpc

import (
	"bytes"
)

type CmdDevelopRequestBotZones struct {
}

func (rsp *CmdDevelopRequestBotZones) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdDevelopRequestBotZones) Serialize(buf *bytes.Buffer) error {
	return nil
}
