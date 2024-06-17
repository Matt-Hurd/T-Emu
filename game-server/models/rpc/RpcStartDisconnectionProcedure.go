package rpc

import (
	"bytes"
	"game-server/helpers"
)

type RpcStartDisconnectionProcedure struct {
	disconnectionCode int32
	additionalInfo    string
	technicalMessage  string
}

func (rsp *RpcStartDisconnectionProcedure) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadInt32(buf, &rsp.disconnectionCode); err != nil {
		return err
	}
	if err := helpers.ReadString(buf, &rsp.additionalInfo); err != nil {
		return err
	}
	if err := helpers.ReadString(buf, &rsp.technicalMessage); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcStartDisconnectionProcedure) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WriteInt32(buf, rsp.disconnectionCode); err != nil {
		return err
	}
	if err := helpers.WriteString(buf, rsp.additionalInfo); err != nil {
		return err
	}
	if err := helpers.WriteString(buf, rsp.technicalMessage); err != nil {
		return err
	}
	return nil
}
