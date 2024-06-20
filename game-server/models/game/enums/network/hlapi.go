package network

import (
	"fmt"
)

type PacketID uint16

const (
	UserMessage PacketID = iota
	ObjectDestroy
	RPC
	ObjectSpawn
	Owner
	Command
	LocalPlayerTransform
	SyncEvent
	UpdateVars
	SyncList
	ObjectSpawnScene
	NetworkInfo
	SpawnFinished
	ObjectHide
	CRC
	LocalClientAuthority
	LocalChildTransform
	Fragment
	PeerClientAuthority
	HLAPIMsg = iota + 9
	LLAPIMsg
	HLAPIResend
	HLAPIPending
	Connect
	Disconnect
	Error
	Ready
	NotReady
	AddPlayer
	RemovePlayer
	Scene
	Animation
	AnimationParameters
	AnimationTrigger
	LobbyReadyToBegin
	LobbySceneLoaded
	LobbyAddPlayerFailed
	LobbyReturnToLobby
	ReconnectPlayer
	ConnectionRequest       PacketID = 147
	RejectResponse          PacketID = 148
	WorldSpawn              PacketID = 151
	WorldUnspawn            PacketID = 152
	SubWorldUnspawn         PacketID = 154
	PlayerSpawn             PacketID = 155
	PlayerUnspawn           PacketID = 156
	ObserverSpawn           PacketID = 157
	ObserverUnspawn         PacketID = 158
	DeathInventorySync      PacketID = 160
	BEPacket                PacketID = 168
	messageFromServer       PacketID = 170
	SpawnObservedPlayer     PacketID = 171
	SpawnObservedPlayers    PacketID = 172
	ChangeFramerate         PacketID = 175
	SnapshotObservedPlayers PacketID = 173
	CommandsObservedPlayers PacketID = 174
	SnapshotBTRVehicles     PacketID = 184
	PartialCommand          PacketID = 185
	NightMare               PacketID = 188
	SyncToPlayers           PacketID = 189
	ProgressReport          PacketID = 190
	SubWorldSpawnLoot       PacketID = 191
	SubWorldSpawnSearchLoot PacketID = 192
	HLAPI                   PacketID = 18385
)

var msgTypeEnum = map[uint16]string{
	0:     "UserMessage",
	1:     "ObjectDestroy",
	2:     "RPC",
	3:     "ObjectSpawn",
	4:     "Owner",
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
	151:   "WorldSpawn",
	152:   "WorldUnspawn",
	154:   "SubWorldUnspawn",
	155:   "PlayerSpawn",
	156:   "PlayerUnspawn",
	157:   "ObserverSpawn",
	158:   "ObserverUnspawn",
	160:   "DeathInventorySync",
	168:   "BEPacket",
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

func GetPacketType(num uint16) string {
	if msgType, ok := msgTypeEnum[num]; ok {
		return msgType + fmt.Sprintf("(%d)", num)
	}
	return fmt.Sprintf("Unknown(%d)", num)
}
