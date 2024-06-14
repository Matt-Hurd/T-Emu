using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace TarkovPacketSer
{
    [JsonConverter(typeof(StringEnumConverter))]
    public enum MsgTypeEnum : short
    {
        Unknown = 0,

        ObjectDestroy = 1,
        RPC = 2,
        ObjectSpawn = 3,
        Command = 5,
        LocalPlayerTransform = 6,
        SyncEvent = 7,
        UpdateVars = 8,
        SyncList = 9,
        ObjectSpawnScene = 10,
        NetworkInfo = 11,
        SpawnFinished = 12,
        ObjectHide = 13,
        CRC = 14,
        LocalClientAuthority = 15,
        LocalChildTransform = 16,
        Fragment = 17,
        PeerClientAuthority = 18,
        HLAPIMsg = 28,
        LLAPIMsg = 29,
        HLAPIResend = 30,
        HLAPIPending = 31,
        Connect = 32,
        Disconnect = 33,
        Error = 34,
        Ready = 35,
        NotReady = 36,
        AddPlayer = 37,
        RemovePlayer = 38,
        Scene = 39,
        Animation = 40,
        AnimationParameters = 41,
        AnimationTrigger = 42,
        LobbyReadyToBegin = 43,
        LobbySceneLoaded = 44,
        LobbyAddPlayerFailed = 45,
        LobbyReturnToLobby = 46,
        ReconnectPlayer = 47,

        ConnectionRequest = 147,
        RejectResponse = 148,
        BEPacket = 168,
        PartialCommand = 185,
        NightMare = 188,
        SyncToPlayers = 189,

        WorldSpawn = 151,
        WorldUnspawn = 152,
        SubWorldSpawnLoot = 191,
        SubWorldSpawnSearchLoot = 192,
        SubWorldUnspawn = 154,
        PlayerUnspawn = 156,
        ObserverUnspawn = 158,
        DeathInventorySync = 160,

        PlayerSpawn = 155,
        ObserverSpawn = 157,

        messageFromServer = 170,
        SpawnObservedPlayer = 171,
        SpawnObservedPlayers = 172,
        ChangeFramerate = 175,
        SnapshotObservedPlayers = 173,
        CommandsObservedPlayers = 174,
        SnapshotBTRVehicles = 184,

        HLAPI = 18385,

        ProgressReport = 190
    }
}