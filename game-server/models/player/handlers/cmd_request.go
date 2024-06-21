package handlers

import (
	networkModels "game-server/models/game/enums/network"
	"game-server/models/game/request"
	"game-server/models/game/response"
	"game-server/models/rpc"
	"game-server/network"
)

func HandlePacketCmdRequest(packet *request.PacketCmdRequest, g *network.NetworkManager) {
	if packet.CmdId == int32(networkModels.CmdSpawn) {
		testPacket := &response.WorldSpawn{}
		err := testPacket.GetDefault()
		if err != nil {
			panic(err)
		}

		g.SendReliableDataPacket(networkModels.WorldSpawn, testPacket)
		g.SendReliableDataPacket(networkModels.RPC, &response.PacketRpcResponse{
			CmdId:   int32(networkModels.RpcGameSpawned),
			NetId:   1,
			Command: &rpc.RpcGameSpawned{},
		})

		g.SendReliableDataPacket(networkModels.RPC, &response.PacketRpcResponse{
			CmdId: int32(networkModels.RpcGameStarting),
			NetId: 1,
			Command: &rpc.RpcGameStarting{
				Seconds: 10,
			}})
	}
}
