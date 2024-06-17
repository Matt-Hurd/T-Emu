package rpc

import (
	"bytes"
)

type CmdDevelopRequestBotProfiles struct {
}

func (rsp *CmdDevelopRequestBotProfiles) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdDevelopRequestBotProfiles) Serialize(buf *bytes.Buffer) error {
	return nil
}
