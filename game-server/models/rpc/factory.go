package rpc

import (
	"bytes"
	"fmt"
	"sync"
)

type RPCCommand interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

var rpcCommandFactory = map[int32]func() RPCCommand{
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

func GetRPCCommand(cmdID int32) (RPCCommand, bool) {
	mu.Lock()
	defer mu.Unlock()

	if factory, ok := rpcCommandFactory[cmdID]; ok {
		return factory(), true
	}
	return nil, false
}

func GetRpcType(num int32) string {
	rpcType := map[int32]string{
		-1723132743: "CmdSpawn",
		740792038:   "CmdRespawn",
		-1220356686: "CmdStartGame",
		1792897173:  "CmdStartGameAfterTeleport",
		273195288:   "CmdRestartGameInitiate",
		-1501005473: "CmdRestartGame",
		-40021267:   "CmdGameStarted",
		-750099178:  "CmdStopGame",
		463608476:   "CmdSyncGameTime",
		-1035840717: "CmdDevelopRequestBot",
		1432950484:  "CmdDevelopRequestBotZones",
		930653927:   "CmdDevelopRequestBotGroups",
		-503414723:  "CmdDevelopRequestBotProfiles",
		-1581543574: "CmdDevelopmentSpawnBotRequest",
		102630535:   "CmdDevelopmentSpawnBotOnServer",
		-349255409:  "CmdDevelopmentSpawnBotOnClient",
		-1733636721: "CmdDisconnectAcceptedOnClient",
		1240699829:  "CmdWorldSpawnConfirm",
		-1317447737: "CmdSpawnConfirm",
		810388720:   "CmdReportVoipAbuse",
		905971479:   "CmdPlayerEffectsPause",
		-65034947:   "CmdOnPlayerKeeperStatisticsChanged",
		-942910572:  "CmdGetRadiotransmitterData",
		1404634890:  "CmdGetTraderServicesData",
		-1952818640: "RpcGameSpawned",
		2117859815:  "RpcGameMatching",
		-1157222870: "RpcGameStarting",
		1572370779:  "RpcGameStartingWithTeleport",
		-1838445225: "RpcGameStarted",
		94275293:    "RpcGameRestarting",
		-1243884988: "RpcGameRestarted",
		-758380962:  "RpcGameStopping",
		-1825579357: "RpcGameStopped",
		547040626:   "RpcSyncGameTime",
		1152897188:  "RpcDevelopSendBotData",
		-1920895376: "RpcDevelopSendBotDataZone",
		314346392:   "RpcDevelopSendBotDataGroups",
		-69469010:   "RpcDevelopSendBotDataProfiles",
		-1269941968: "RpcDevelopmentSpawnBotResponse",
		-435294673:  "RpcSoftStopNotification",
		1124901489:  "RpcStartDisconnectionProcedure",
		1547608889:  "RpcVoipAbuseNotification",
		-2040405782: "RpcAirdropContainerData",
		-689857055:  "RpcMineDirectionExplosion",
		-2141949542: "RpcSuccessAirdropFlareEvent",
		778150830:   "RpcBufferZoneData",
		-52162261:   "RpcSendClientRadioTransmitterData",
		1358208182:  "RpcSendObserverRadioTransmitterData",
		-1785644202: "RpcSendTraderServicesData",
		361141025:   "RpcSyncLighthouseTraderZoneData",
		-1536532007: "RpcSendCompletedAchievementsData",
	}
	if msgType, ok := rpcType[num]; ok {
		return msgType + fmt.Sprintf("(%d)", num)
	}
	return fmt.Sprintf("Unknown(%d)", num)
}

// {'CmdSpawn': 'fbb9144b99',
//  'CmdRespawn': 'fbe696272c',
//  'CmdStartGame': 'fbb2d542b7',
//  'CmdStartGameAfterTeleport': 'fb9570dd6a',
//  'CmdRestartGameInitiate': 'fb18a14810',
//  'CmdRestartGame': 'fb5f7988a6',
//  'CmdGameStarted': 'fbed529dfd',
//  'CmdStopGame': 'fb16654ad3',
//  'CmdSyncGameTime': 'fb9c1aa21b',
//  'CmdDevelopRequestBot': 'fb335342c2',
//  'CmdDevelopRequestBotZones': 'fbd4166955',
//  'CmdDevelopRequestBotGroups': 'fbe7a67837',
//  'CmdDevelopRequestBotProfiles': 'fb3d80fee1',
//  'CmdDevelopmentSpawnBotRequest': 'fb6a8fbba1',
//  'CmdDevelopmentSpawnBotOnServer': 'fb87041e06',
//  'CmdDevelopmentSpawnBotOnClient': 'fb0fc92eeb',
//  'CmdDisconnectAcceptedOnClient': 'fb8fcdaa98',
//  'CmdWorldSpawnConfirm': 'fbb593f349',
//  'CmdSpawnConfirm': 'fbc75779b1',
//  'CmdReportVoipAbuse': 'fbf08c4d30',
//  'CmdPlayerEffectsPause': 'fb17070036',
//  'CmdOnPlayerKeeperStatisticsChanged': 'fb3da51ffc',
//  'CmdGetRadiotransmitterData': 'fb9453ccc7',
//  'CmdGetTraderServicesData': 'fb0a07b953',
//  'RpcGameSpawned': 'fb305a9a8b',
//  'RpcGameMatching': 'fbe7f93b7e',
//  'RpcGameStarting': 'fb2a2e06bb',
//  'RpcGameStartingWithTeleport': 'fb5b79b85d',
//  'RpcGameStarted': 'fb578d6b92',
//  'RpcGameRestarting': 'fbdd869e05',
//  'RpcGameRestarted': 'fb44d2dbb5',
//  'RpcGameStopping': 'fb5e06ccd2',
//  'RpcGameStopped': 'fba3de2f93',
//  'RpcSyncGameTime': 'fb722d9b20',
//  'RpcDevelopSendBotData': 'fba4d0b744',
//  'RpcDevelopSendBotDataZone': 'fb7076818d',
//  'RpcDevelopSendBotDataGroups': 'fb988bbc12',
//  'RpcDevelopSendBotDataProfiles': 'fbaefcdbfb',
//  'RpcDevelopmentSpawnBotResponse': 'fb30394eb4',
//  'RpcSoftStopNotification': 'fb2fee0de6',
//  'RpcStartDisconnectionProcedure': 'fb71a20c43',
//  'RpcVoipAbuseNotification': 'fb39a33e5c',
//  'RpcAirdropContainerData': 'fbeae06186',
//  'RpcMineDirectionExplosion': 'fbe19de1d6',
//  'RpcSuccessAirdropFlareEvent': 'fb9a715480',
//  'RpcBufferZoneData': 'fbaea3612e',
//  'RpcSendClientRadioTransmitterData': 'fb2b11e4fc',
//  'RpcSendObserverRadioTransmitterData': 'fbb69cf450',
//  'RpcSendTraderServicesData': 'fb563b9195',
//  'RpcSyncLighthouseTraderZoneData': 'fb21938615',
//  'RpcSendCompletedAchievementsData': 'fbd9616aa4'}
