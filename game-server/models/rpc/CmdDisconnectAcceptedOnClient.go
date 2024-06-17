package rpc

import (
	"bytes"
)

type CmdDisconnectAcceptedOnClient struct {
}

func (rsp *CmdDisconnectAcceptedOnClient) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdDisconnectAcceptedOnClient) Serialize(buf *bytes.Buffer) error {
	return nil
}
