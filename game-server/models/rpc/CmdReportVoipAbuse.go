package rpc

import (
	"bytes"
)

type CmdReportVoipAbuse struct {
}

func (rsp *CmdReportVoipAbuse) Deserialize(buf *bytes.Buffer) error {
	return nil
}

func (rsp *CmdReportVoipAbuse) Serialize(buf *bytes.Buffer) error {
	return nil
}
