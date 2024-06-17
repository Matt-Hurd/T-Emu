package rpc

import (
	"bytes"
)

type RpcSendObserverRadioTransmitterData struct {
	data GStruct134
}

func (rsp *RpcSendObserverRadioTransmitterData) Deserialize(buf *bytes.Buffer) error {
	if err := rsp.data.Deserialize(buf); err != nil {
		return err
	}
	return nil
}

func (rsp *RpcSendObserverRadioTransmitterData) Serialize(buf *bytes.Buffer) error {
	if err := rsp.data.Serialize(buf); err != nil {
		return err
	}
	return nil
}
