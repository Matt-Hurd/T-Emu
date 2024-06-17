package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/rpc"
)

type PacketRpcResponse struct {
	CmdId   uint32
	NetId   uint32
	Command rpc.RPCCommand
}

func (p *PacketRpcResponse) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadPackedUInt32(buffer, &p.CmdId)
	if err != nil {
		panic(err)
	}

	err = helpers.ReadPackedUInt32(buffer, &p.NetId)
	if err != nil {
		panic(err)
	}

	rpcCmd, success := rpc.GetRPCCommand(int(p.CmdId))
	if !success {
		panic("Failed to get rpc command")
	}
	p.Command = rpcCmd
	p.Command.Deserialize(buffer)
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
