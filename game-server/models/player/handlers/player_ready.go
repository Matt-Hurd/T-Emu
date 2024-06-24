package handlers

import (
	"game-server/helpers"
	networkModels "game-server/models/game/enums/network"
	"game-server/models/game/math"
	"game-server/models/game/request"
	"game-server/models/game/response"
	"game-server/models/rpc"
	"game-server/network"
	"time"

	"github.com/g3n/engine/math32"
)

func HandlePacketPlayerReady(packet *request.PacketClientReady, g *network.NetworkManager) {
	g.SendReliableDataPacket(networkModels.SpawnFinished, &response.PacketSpawnFinished{State: 0})
	g.SendReliableDataPacket(networkModels.SpawnFinished, &response.PacketSpawnFinished{State: 1})

	objectSpawnPacket := response.PacketObjectSpawn{
		NetId:    1,
		Position: math.Vector3{Vector3: math32.Vector3{X: 0, Y: 0, Z: 0}},
		Payload:  []byte{},
		Rotation: math.Quaternion{
			Quaternion: math32.Quaternion{X: 0, Y: 0, Z: 0, W: 1}},
	}
	objectSpawnPacket.AssetId.FromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd})
	g.SendReliableDataPacket(networkModels.ObjectSpawn, &objectSpawnPacket)
	g.SendReliableDataPacket(networkModels.LocalClientAuthority, &response.PacketClientAuthority{NetId: 1, Authority: true})
	g.SendReliableDataPacket(networkModels.RPC, &response.PacketRpcResponse{
		CmdId: int32(networkModels.RpcSyncGameTime),
		NetId: 1,
		Command: &rpc.RpcSyncGameTime{
			Time: uint64(helpers.TimeToInt64(time.Now().Add(20 * time.Second))),
		},
	})
}
