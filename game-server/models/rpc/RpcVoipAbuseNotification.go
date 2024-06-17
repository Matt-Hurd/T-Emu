package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcVoipAbuseNotification struct {
	reporter string
}

func (rsp *RpcVoipAbuseNotification) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadString(buf, &rsp.reporter); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcVoipAbuseNotification) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteString(buf, rsp.reporter); err != nil {
		return err
	}
	return nil
}
