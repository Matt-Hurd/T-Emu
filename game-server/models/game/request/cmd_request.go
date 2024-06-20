package request

import (
	"bytes"
	"fmt"
	"game-server/helpers"
	"game-server/models/rpc"
)

type PacketCmdRequest struct {
	CmdId   int32
	NetId   uint32
	Command rpc.RPCCommand
}

func (p *PacketCmdRequest) Deserialize(buffer *bytes.Buffer) error {
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

	fmt.Printf("Deserialized PacketCmdRequest: %v\n", p)
	rpcCmd, success := rpc.GetRPCCommand(int32(p.CmdId))
	if !success {
		panic("Failed to get rpc command")
	}
	p.Command = rpcCmd
	p.Command.Deserialize(buffer)

	fmt.Printf("Deserialized PacketCmdRequest: %v\n", p)
	return nil
}

func (p *PacketCmdRequest) Serialize(buffer *bytes.Buffer) error {
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
