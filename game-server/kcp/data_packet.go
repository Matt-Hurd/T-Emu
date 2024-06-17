package kcp

import (
	"fmt"
	"game-server/helpers"
	"game-server/models"
	"game-server/models/game/request"
	"game-server/models/game/response"
	"game-server/models/rpc"
)

func HandleDataPacket(packet *models.GClass2498, kcp_0 *GClass2486) {
	dataPacket := &models.DataPacket{}
	err := dataPacket.Parse(packet.Buffer)
	if err != nil {
		fmt.Println("Error parsing data packet:", err)
		return
	}
	switch dataPacket.GamePacketType {
	case 5:
		HandlePacketCmdRequest(dataPacket.GamePacket.(*request.PacketCmdRequest), kcp_0)
	case 147:
		HandlePacketConnection(dataPacket.GamePacket.(*request.PacketConnection), kcp_0)
	case 190:
		HandlePacketProgressReport(dataPacket.GamePacket.(*request.PacketProgressReport))
	default:
		HandleUnknownPacket(dataPacket, packet)
	}
}

func HandlePacketCmdRequest(packet *request.PacketCmdRequest, g *GClass2486) {
	fmt.Printf("Received PacketCmdRequest: %v\n", packet)
}

func HandlePacketProgressReport(packet *request.PacketProgressReport) {
	fmt.Printf("Received PacketProgressReport: %v\n", packet)
}

func HandlePacketConnection(packet *request.PacketConnection, g *GClass2486) {
	var err error

	connectionResponse := &response.PacketConnection{}
	resp := &models.DataPacket{
		GamePacketType: 147,
		GamePacket:     connectionResponse,
	}
	connectionResponse.GetDefault()
	g.SendQueue <- &models.GClass2498{
		Channel: models.NetworkChannelReliable,
		Type:    models.NetworkMessageTypeData,
		Buffer:  resp.Write(),
	}

	nightmare := &response.NightMare{
		Id: 0,
	}
	resp = &models.DataPacket{
		GamePacketType: 188,
		GamePacket:     nightmare,
	}
	nightmare.PrefabsData, err = helpers.CompressZlib([]byte("[]"))
	if err != nil {
		fmt.Println("Error compressing prefabs data:", err)
	}
	nightmare.CustomizationData, err = helpers.CompressZlib([]byte(`["66043cc27502eca33a08cad0","5e9dc97c86f774054c19ac9a"]`))
	if err != nil {
		fmt.Println("Error compressing customization data:", err)
	}
	g.SendQueue <- &models.GClass2498{
		Channel: models.NetworkChannelReliable,
		Type:    models.NetworkMessageTypeData,
		Buffer:  resp.Write(),
	}

	rpcCmd := &response.PacketRpcResponse{
		CmdId: 547040626,
		NetId: 1,
		Command: &rpc.RpcSyncGameTime{
			Time: 5250227001791798016,
		},
	}

	resp = &models.DataPacket{
		GamePacketType: 2,
		GamePacket:     rpcCmd,
	}
	g.SendQueue <- &models.GClass2498{
		Channel: models.NetworkChannelReliable,
		Type:    models.NetworkMessageTypeData,
		Buffer:  resp.Write(),
	}
}

func HandleUnknownPacket(packet *models.DataPacket, packet_1 *models.GClass2498) {
	fmt.Printf("Unhandled Packet Type: %s, Reliable(%t), Data: %v\n", GetPacketType(packet), packet_1.Channel == models.NetworkChannelReliable, packet_1.Buffer[4:])
}

func GetPacketType(packet *models.DataPacket) string {
	msgTypeEnum := map[int]string{
		0:     "Unknown",
		1:     "ObjectDestroy",
		2:     "RPC",
		3:     "ObjectSpawn",
		5:     "Command",
		6:     "LocalPlayerTransform",
		7:     "SyncEvent",
		8:     "UpdateVars",
		9:     "SyncList",
		10:    "ObjectSpawnScene",
		11:    "NetworkInfo",
		12:    "SpawnFinished",
		13:    "ObjectHide",
		14:    "CRC",
		15:    "LocalClientAuthority",
		16:    "LocalChildTransform",
		17:    "Fragment",
		18:    "PeerClientAuthority",
		28:    "HLAPIMsg",
		29:    "LLAPIMsg",
		30:    "HLAPIResend",
		31:    "HLAPIPending",
		32:    "Connect",
		33:    "Disconnect",
		34:    "Error",
		35:    "Ready",
		36:    "NotReady",
		37:    "AddPlayer",
		38:    "RemovePlayer",
		39:    "Scene",
		40:    "Animation",
		41:    "AnimationParameters",
		42:    "AnimationTrigger",
		43:    "LobbyReadyToBegin",
		44:    "LobbySceneLoaded",
		45:    "LobbyAddPlayerFailed",
		46:    "LobbyReturnToLobby",
		47:    "ReconnectPlayer",
		147:   "ConnectionRequest",
		148:   "RejectResponse",
		168:   "BEPacket",
		151:   "WorldSpawn",
		152:   "WorldUnspawn",
		154:   "SubWorldUnspawn",
		155:   "PlayerSpawn",
		156:   "PlayerUnspawn",
		157:   "ObserverSpawn",
		158:   "ObserverUnspawn",
		160:   "DeathInventorySync",
		170:   "messageFromServer",
		171:   "SpawnObservedPlayer",
		172:   "SpawnObservedPlayers",
		175:   "ChangeFramerate",
		173:   "SnapshotObservedPlayers",
		174:   "CommandsObservedPlayers",
		184:   "SnapshotBTRVehicles",
		185:   "PartialCommand",
		188:   "NightMare",
		189:   "SyncToPlayers",
		190:   "ProgressReport",
		191:   "SubWorldSpawnLoot",
		192:   "SubWorldSpawnSearchLoot",
		18385: "HLAPI",
	}
	if msgType, ok := msgTypeEnum[int(packet.GamePacketType)]; ok {
		return msgType + fmt.Sprintf("(%d)", packet.GamePacketType)
	}
	return fmt.Sprintf("Unknown(%d)", packet.GamePacketType)
}
