package network

type RpcID int32

const (
	CmdSpawn                            RpcID = -1723132743
	CmdRespawn                          RpcID = 740792038
	CmdStartGame                        RpcID = -1220356686
	CmdStartGameAfterTeleport           RpcID = 1792897173
	CmdRestartGameInitiate              RpcID = 273195288
	CmdRestartGame                      RpcID = -1501005473
	CmdGameStarted                      RpcID = -40021267
	CmdStopGame                         RpcID = -750099178
	CmdSyncGameTime                     RpcID = 463608476
	CmdDevelopRequestBot                RpcID = -1035840717
	CmdDevelopRequestBotZones           RpcID = 1432950484
	CmdDevelopRequestBotGroups          RpcID = 930653927
	CmdDevelopRequestBotProfiles        RpcID = -503414723
	CmdDevelopmentSpawnBotRequest       RpcID = -1581543574
	CmdDevelopmentSpawnBotOnServer      RpcID = 102630535
	CmdDevelopmentSpawnBotOnClient      RpcID = -349255409
	CmdDisconnectAcceptedOnClient       RpcID = -1733636721
	CmdWorldSpawnConfirm                RpcID = 1240699829
	CmdSpawnConfirm                     RpcID = -1317447737
	CmdReportVoipAbuse                  RpcID = 810388720
	CmdPlayerEffectsPause               RpcID = 905971479
	CmdOnPlayerKeeperStatisticsChanged  RpcID = -65034947
	CmdGetRadiotransmitterData          RpcID = -942910572
	CmdGetTraderServicesData            RpcID = 1404634890
	RpcGameSpawned                      RpcID = -1952818640
	RpcGameMatching                     RpcID = 2117859815
	RpcGameStarting                     RpcID = -1157222870
	RpcGameStartingWithTeleport         RpcID = 1572370779
	RpcGameStarted                      RpcID = -1838445225
	RpcGameRestarting                   RpcID = 94275293
	RpcGameRestarted                    RpcID = -1243884988
	RpcGameStopping                     RpcID = -758380962
	RpcGameStopped                      RpcID = -1825579357
	RpcSyncGameTime                     RpcID = 547040626
	RpcDevelopSendBotData               RpcID = 1152897188
	RpcDevelopSendBotDataZone           RpcID = -1920895376
	RpcDevelopSendBotDataGroups         RpcID = 314346392
	RpcDevelopSendBotDataProfiles       RpcID = -69469010
	RpcDevelopmentSpawnBotResponse      RpcID = -1269941968
	RpcSoftStopNotification             RpcID = -435294673
	RpcStartDisconnectionProcedure      RpcID = 1124901489
	RpcVoipAbuseNotification            RpcID = 1547608889
	RpcAirdropContainerData             RpcID = -2040405782
	RpcMineDirectionExplosion           RpcID = -689857055
	RpcSuccessAirdropFlareEvent         RpcID = -2141949542
	RpcBufferZoneData                   RpcID = 778150830
	RpcSendClientRadioTransmitterData   RpcID = -52162261
	RpcSendObserverRadioTransmitterData RpcID = 1358208182
	RpcSendTraderServicesData           RpcID = -1785644202
	RpcSyncLighthouseTraderZoneData     RpcID = 361141025
	RpcSendCompletedAchievementsData    RpcID = -1536532007
)

var rpcIDToString = map[RpcID]string{
	CmdSpawn:                            "CmdSpawn",
	CmdRespawn:                          "CmdRespawn",
	CmdStartGame:                        "CmdStartGame",
	CmdStartGameAfterTeleport:           "CmdStartGameAfterTeleport",
	CmdRestartGameInitiate:              "CmdRestartGameInitiate",
	CmdRestartGame:                      "CmdRestartGame",
	CmdGameStarted:                      "CmdGameStarted",
	CmdStopGame:                         "CmdStopGame",
	CmdSyncGameTime:                     "CmdSyncGameTime",
	CmdDevelopRequestBot:                "CmdDevelopRequestBot",
	CmdDevelopRequestBotZones:           "CmdDevelopRequestBotZones",
	CmdDevelopRequestBotGroups:          "CmdDevelopRequestBotGroups",
	CmdDevelopRequestBotProfiles:        "CmdDevelopRequestBotProfiles",
	CmdDevelopmentSpawnBotRequest:       "CmdDevelopmentSpawnBotRequest",
	CmdDevelopmentSpawnBotOnServer:      "CmdDevelopmentSpawnBotOnServer",
	CmdDevelopmentSpawnBotOnClient:      "CmdDevelopmentSpawnBotOnClient",
	CmdDisconnectAcceptedOnClient:       "CmdDisconnectAcceptedOnClient",
	CmdWorldSpawnConfirm:                "CmdWorldSpawnConfirm",
	CmdSpawnConfirm:                     "CmdSpawnConfirm",
	CmdReportVoipAbuse:                  "CmdReportVoipAbuse",
	CmdPlayerEffectsPause:               "CmdPlayerEffectsPause",
	CmdOnPlayerKeeperStatisticsChanged:  "CmdOnPlayerKeeperStatisticsChanged",
	CmdGetRadiotransmitterData:          "CmdGetRadiotransmitterData",
	CmdGetTraderServicesData:            "CmdGetTraderServicesData",
	RpcGameSpawned:                      "RpcGameSpawned",
	RpcGameMatching:                     "RpcGameMatching",
	RpcGameStarting:                     "RpcGameStarting",
	RpcGameStartingWithTeleport:         "RpcGameStartingWithTeleport",
	RpcGameStarted:                      "RpcGameStarted",
	RpcGameRestarting:                   "RpcGameRestarting",
	RpcGameRestarted:                    "RpcGameRestarted",
	RpcGameStopping:                     "RpcGameStopping",
	RpcGameStopped:                      "RpcGameStopped",
	RpcSyncGameTime:                     "RpcSyncGameTime",
	RpcDevelopSendBotData:               "RpcDevelopSendBotData",
	RpcDevelopSendBotDataZone:           "RpcDevelopSendBotDataZone",
	RpcDevelopSendBotDataGroups:         "RpcDevelopSendBotDataGroups",
	RpcDevelopSendBotDataProfiles:       "RpcDevelopSendBotDataProfiles",
	RpcDevelopmentSpawnBotResponse:      "RpcDevelopmentSpawnBotResponse",
	RpcSoftStopNotification:             "RpcSoftStopNotification",
	RpcStartDisconnectionProcedure:      "RpcStartDisconnectionProcedure",
	RpcVoipAbuseNotification:            "RpcVoipAbuseNotification",
	RpcAirdropContainerData:             "RpcAirdropContainerData",
	RpcMineDirectionExplosion:           "RpcMineDirectionExplosion",
	RpcSuccessAirdropFlareEvent:         "RpcSuccessAirdropFlareEvent",
	RpcBufferZoneData:                   "RpcBufferZoneData",
	RpcSendClientRadioTransmitterData:   "RpcSendClientRadioTransmitterData",
	RpcSendObserverRadioTransmitterData: "RpcSendObserverRadioTransmitterData",
	RpcSendTraderServicesData:           "RpcSendTraderServicesData",
	RpcSyncLighthouseTraderZoneData:     "RpcSyncLighthouseTraderZoneData",
	RpcSendCompletedAchievementsData:    "RpcSendCompletedAchievementsData",
}

func (rpcID RpcID) String() string {
	return rpcIDToString[rpcID]
}
