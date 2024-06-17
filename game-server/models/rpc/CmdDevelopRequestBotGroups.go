package rpc

import (
	"bytes"
)

type CmdDevelopRequestBotGroups struct {
}

func (rsp *CmdDevelopRequestBotGroups) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdDevelopRequestBotGroups) Serialize(buf *bytes.Buffer) error {
	return nil
}
