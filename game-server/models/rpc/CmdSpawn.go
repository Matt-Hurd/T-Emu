package rpc

import (
	"bytes"
)

type CmdSpawn struct {
}

func (rsp *CmdSpawn) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdSpawn) Serialize(buf *bytes.Buffer) error {
	return nil
}
