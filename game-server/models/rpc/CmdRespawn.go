package rpc

import (
	"bytes"
)

type CmdRespawn struct {
}

func (rsp *CmdRespawn) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdRespawn) Serialize(buf *bytes.Buffer) error {
	return nil
}
