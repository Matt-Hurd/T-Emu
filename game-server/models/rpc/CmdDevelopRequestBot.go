package rpc

import (
	"bytes"
	"game-server/helpers"
)

type CmdDevelopRequestBot struct {
	profileId string
}

func (rsp *CmdDevelopRequestBot) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.profileId); err != nil {
		return err
	}
	return nil
}

func (rsp *CmdDevelopRequestBot) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.profileId); err != nil {
		return err
	}
	return nil
}
