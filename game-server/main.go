package main

import (
	"bytes"
	"fmt"
	"game-server/helpers"
	"game-server/models/rpc"
)

func main() {
	test()
	// srv := server.NewServer()
	// err := srv.Start()
	// if err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}

func test() {
	// testPacket := &game.PacketConnection{}
	// testPacket.GetDefault()
	// testPacket.Deserialize(bytes.NewBuffer(bigPacket.Buffer[4:]))
	// testPacket.DebugPrintAll()
	// buf := bytes.Buffer{}
	// testPacket.Serialize(&buf)

	// secondPacket := &game.PacketConnection{}
	// secondPacket.Deserialize(&buf)

	// secondPacket.DebugPrintAll()

	bigData := []byte{0xFB, 0xB9, 0x14, 0x4B, 0x99, 0x01}
	buffer := bytes.NewBuffer(bigData)
	var commandId uint32
	var netId uint32
	err := helpers.ReadPackedUInt32(buffer, &commandId)
	if err != nil {
		panic(err)
	}
	err = helpers.ReadPackedUInt32(buffer, &netId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("CommandId: %d\n", commandId)
	fmt.Printf("NetId: %d\n", netId)

	rpcCmd, success := rpc.GetRPCCommand(int(commandId))
	if !success {
		fmt.Printf("Failed to get rpc command with id: %d\n", commandId)
		panic("Failed to get rpc command")
	}
	rpcCmd.Deserialize(buffer)
	fmt.Printf("%+v\n", rpcCmd)
}
