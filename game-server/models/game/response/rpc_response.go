package response

import (
	"bytes"
	"fmt"
	"game-server/helpers"
	"game-server/models/rpc"
)

type PacketRpcResponse struct {
	CmdId   int32
	NetId   uint32
	Command rpc.RPCCommand
}

func (p *PacketRpcResponse) Deserialize(buffer *bytes.Buffer) error {
	var tmp uint32
	err := helpers.ReadPackedUInt32(buffer, &tmp)
	if err != nil {
		panic(err)
	}
	p.CmdId = int32(tmp)

	err = helpers.ReadPackedUInt32(buffer, &p.NetId)
	if err != nil {
		panic(err)
	}

	rpcCmd, success := rpc.GetRPCCommand(int32(p.CmdId))
	if !success {
		panic("Failed to get rpc command")
	}
	p.Command = rpcCmd
	p.Command.Deserialize(buffer)

	fmt.Printf("Deserialized PacketRpcResponse (%s): %v\n", rpc.GetRpcType(p.CmdId), p)
	return nil
}

func (p *PacketRpcResponse) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WritePackedUInt32(buffer, uint32(p.CmdId))
	if err != nil {
		return err
	}
	err = helpers.WritePackedUInt32(buffer, uint32(p.NetId))
	if err != nil {
		return err
	}
	err = p.Command.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}
