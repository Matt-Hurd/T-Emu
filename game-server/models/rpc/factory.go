package rpc

import (
	"bytes"
	"sync"
)

type RPCCommand interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

var rpcCommandFactory = map[int]func() RPCCommand{
	// Commands
	-1723132743: func() RPCCommand { return &CmdSpawn{} },
	740792038:   func() RPCCommand { return &CmdRespawn{} },
	-1220356686: func() RPCCommand { return &CmdStartGame{} },
	1792897173:  func() RPCCommand { return &CmdStartGameAfterTeleport{} },
	273195288:   func() RPCCommand { return &CmdRestartGameInitiate{} },
	-1501005473: func() RPCCommand { return &CmdRestartGame{} },
	-40021267:   func() RPCCommand { return &CmdGameStarted{} },
	-750099178:  func() RPCCommand { return &CmdStopGame{} },
	463608476:   func() RPCCommand { return &CmdSyncGameTime{} },
	-1035840717: func() RPCCommand { return &CmdDevelopRequestBot{} },
	1432950484:  func() RPCCommand { return &CmdDevelopRequestBotZones{} },
	930653927:   func() RPCCommand { return &CmdDevelopRequestBotGroups{} },
	-503414723:  func() RPCCommand { return &CmdDevelopRequestBotProfiles{} },
	-1581543574: func() RPCCommand { return &CmdDevelopmentSpawnBotRequest{} },
	102630535:   func() RPCCommand { return &CmdDevelopmentSpawnBotOnServer{} },
	-349255409:  func() RPCCommand { return &CmdDevelopmentSpawnBotOnClient{} },
	-1733636721: func() RPCCommand { return &CmdDisconnectAcceptedOnClient{} },
	1240699829:  func() RPCCommand { return &CmdWorldSpawnConfirm{} },
	-1317447737: func() RPCCommand { return &CmdSpawnConfirm{} },
	810388720:   func() RPCCommand { return &CmdReportVoipAbuse{} },
	905971479:   func() RPCCommand { return &CmdPlayerEffectsPause{} },
	-65034947:   func() RPCCommand { return &CmdOnPlayerKeeperStatisticsChanged{} },
	-942910572:  func() RPCCommand { return &CmdGetRadiotransmitterData{} },
	1404634890:  func() RPCCommand { return &CmdGetTraderServicesData{} },

	// RPC
	-1952818640: func() RPCCommand { return &RpcGameSpawned{} },
	2117859815:  func() RPCCommand { return &RpcGameMatching{} },
	-1157222870: func() RPCCommand { return &RpcGameStarting{} },
	1572370779:  func() RPCCommand { return &RpcGameStartingWithTeleport{} },
	-1838445225: func() RPCCommand { return &RpcGameStarted{} },
	94275293:    func() RPCCommand { return &RpcGameRestarting{} },
	-1243884988: func() RPCCommand { return &RpcGameRestarted{} },
	-758380962:  func() RPCCommand { return &RpcGameStopping{} },
	-1825579357: func() RPCCommand { return &RpcGameStopped{} },
	547040626:   func() RPCCommand { return &RpcSyncGameTime{} },
	1152897188:  func() RPCCommand { return &RpcDevelopSendBotData{} },
	-1920895376: func() RPCCommand { return &RpcDevelopSendBotDataZone{} },
	314346392:   func() RPCCommand { return &RpcDevelopSendBotDataGroups{} },
	-69469010:   func() RPCCommand { return &RpcDevelopSendBotDataProfiles{} },
	-1269941968: func() RPCCommand { return &RpcDevelopmentSpawnBotResponse{} },
	-435294673:  func() RPCCommand { return &RpcSoftStopNotification{} },
	1124901489:  func() RPCCommand { return &RpcStartDisconnectionProcedure{} },
	1547608889:  func() RPCCommand { return &RpcVoipAbuseNotification{} },
	-2040405782: func() RPCCommand { return &RpcAirdropContainerData{} },
	-689857055:  func() RPCCommand { return &RpcMineDirectionExplosion{} },
	-2141949542: func() RPCCommand { return &RpcSuccessAirdropFlareEvent{} },
	778150830:   func() RPCCommand { return &RpcBufferZoneData{} },
	-52162261:   func() RPCCommand { return &RpcSendClientRadioTransmitterData{} },
	1358208182:  func() RPCCommand { return &RpcSendObserverRadioTransmitterData{} },
	-1785644202: func() RPCCommand { return &RpcSendTraderServicesData{} },
	361141025:   func() RPCCommand { return &RpcSyncLighthouseTraderZoneData{} },
	-1536532007: func() RPCCommand { return &RpcSendCompletedAchievementsData{} },
}

var mu sync.Mutex

func GetRPCCommand(cmdID int) (RPCCommand, bool) {
	mu.Lock()
	defer mu.Unlock()

	if factory, ok := rpcCommandFactory[cmdID]; ok {
		return factory(), true
	}
	return nil, false
}
