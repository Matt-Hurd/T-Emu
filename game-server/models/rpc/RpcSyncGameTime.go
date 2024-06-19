package rpc

import (
	"bytes"
	"fmt"
	"game-server/helpers"
)

type RpcSyncGameTime struct {
	Time uint64
}

func (rsp *RpcSyncGameTime) Deserialize(buf *bytes.Buffer) error {
	if err := helpers.ReadPackedUInt64(buf, &rsp.Time); err != nil {
		return err
	}
	fmt.Printf("Time: %v\n", helpers.Int64ToTime(int64(rsp.Time)))
	return nil
}

func (rsp *RpcSyncGameTime) Serialize(buf *bytes.Buffer) error {
	if err := helpers.WritePackedUInt64(buf, rsp.Time); err != nil {
		return err
	}
	return nil
}
