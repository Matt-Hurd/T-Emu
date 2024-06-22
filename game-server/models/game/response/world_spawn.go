package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"game-server/models/game/enums"
)

type WorldSpawn struct {
	ExfiltrationController    core.ExfiltrationController
	BufferZoneControllerClass core.BufferZoneControllerClass
	SmokeGrenades             []core.SmokeGrenadeInfo
	DoorInfos                 []core.DoorInfo
	LampControllerInfos       []core.LampControllerInfo
	WindowBreakerInfos        []core.WindowBreakerInfo
	SynchronizableObjectTypes []SynchronizableObjectType
	BTR                       core.BTR
}

type SynchronizableObjectType struct {
	Type  enums.SynchronizableObjectType
	Value interface {
		Serialize(buffer *bytes.Buffer) error
		Deserialize(buffer *bytes.Buffer) error
	}
}

func (ws *WorldSpawn) Deserialize(buffer *bytes.Buffer) error {
	ws.ExfiltrationController = core.ExfiltrationController{}
	if err := ws.ExfiltrationController.Deserialize(buffer); err != nil {
		return err
	}

	if err := ws.BufferZoneControllerClass.Deserialize(buffer); err != nil {
		return err
	}

	if err := ws.InitializeSmokeGrenades(buffer); err != nil {
		return err
	}

	if err := ws.InitializeDoors(buffer); err != nil {
		return err
	}

	if err := ws.InitializeLampControllers(buffer); err != nil {
		return err
	}

	if err := ws.InitializeWindowBreakers(buffer); err != nil {
		return err
	}

	if err := ws.SynchronizableObjectType(buffer); err != nil {
		return err
	}

	if err := ws.ReadBTR(buffer); err != nil {
		return err
	}

	return nil
}

func (ws *WorldSpawn) Serialize(buffer *bytes.Buffer) error {
	if err := ws.ExfiltrationController.Serialize(buffer); err != nil {
		return err
	}

	if err := ws.BufferZoneControllerClass.Serialize(buffer); err != nil {
		return err
	}

	if err := ws.SerializeSmokeGrenades(buffer); err != nil {
		return err
	}

	if err := ws.SerializeDoors(buffer); err != nil {
		return err
	}

	if err := ws.SerializeLampControllers(buffer); err != nil {
		return err
	}

	if err := ws.SerializeWindowBreakers(buffer); err != nil {
		return err
	}

	if err := ws.SerializeSynchronizableObjectType(buffer); err != nil {
		return err
	}

	if err := ws.BTR.Serialize(buffer); err != nil {
		return err
	}

	return nil
}

func (ws *WorldSpawn) InitializeSmokeGrenades(buffer *bytes.Buffer) error {
	var smokeGrenadesCount int32
	err := helpers.ReadInt32(buffer, &smokeGrenadesCount)
	if err != nil {
		return err
	}

	ws.SmokeGrenades = make([]core.SmokeGrenadeInfo, smokeGrenadesCount)
	for i := 0; i < int(smokeGrenadesCount); i++ {
		var smokeGrenadeInfo core.SmokeGrenadeInfo
		if err := smokeGrenadeInfo.Deserialize(buffer); err != nil {
			return err
		}
		ws.SmokeGrenades[i] = smokeGrenadeInfo
	}
	return nil
}

func (ws *WorldSpawn) SerializeSmokeGrenades(buffer *bytes.Buffer) error {
	smokeGrenadesCount := int32(len(ws.SmokeGrenades))
	if err := helpers.WriteInt32(buffer, smokeGrenadesCount); err != nil {
		return err
	}
	for _, smokeGrenade := range ws.SmokeGrenades {
		if err := smokeGrenade.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ws *WorldSpawn) InitializeDoors(buffer *bytes.Buffer) error {
	var doorsCount int32
	err := helpers.ReadInt32(buffer, &doorsCount)
	if err != nil {
		return err
	}

	ws.DoorInfos = make([]core.DoorInfo, doorsCount)
	for i := 0; i < int(doorsCount); i++ {
		var doorInfo core.DoorInfo
		if err := doorInfo.Deserialize(buffer); err != nil {
			return err
		}
		ws.DoorInfos[i] = doorInfo
	}
	return nil
}

func (ws *WorldSpawn) SerializeDoors(buffer *bytes.Buffer) error {
	doorsCount := int32(len(ws.DoorInfos))
	if err := helpers.WriteInt32(buffer, doorsCount); err != nil {
		return err
	}
	for _, doorInfo := range ws.DoorInfos {
		if err := doorInfo.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ws *WorldSpawn) InitializeLampControllers(buffer *bytes.Buffer) error {
	var lampControllersCount int32
	err := helpers.ReadInt32(buffer, &lampControllersCount)
	if err != nil {
		return err
	}

	ws.LampControllerInfos = make([]core.LampControllerInfo, lampControllersCount)
	for i := 0; i < int(lampControllersCount); i++ {
		var lampControllerInfo core.LampControllerInfo
		if err := lampControllerInfo.Deserialize(buffer); err != nil {
			return err
		}
		ws.LampControllerInfos[i] = lampControllerInfo
	}
	return nil
}

func (ws *WorldSpawn) SerializeLampControllers(buffer *bytes.Buffer) error {
	lampControllersCount := int32(len(ws.LampControllerInfos))
	if err := helpers.WriteInt32(buffer, lampControllersCount); err != nil {
		return err
	}
	for _, lampControllerInfo := range ws.LampControllerInfos {
		if err := lampControllerInfo.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ws *WorldSpawn) InitializeWindowBreakers(buffer *bytes.Buffer) error {
	var windowBreakersCount int32
	err := helpers.ReadInt32(buffer, &windowBreakersCount)
	if err != nil {
		return err
	}

	ws.WindowBreakerInfos = make([]core.WindowBreakerInfo, windowBreakersCount)
	for i := 0; i < int(windowBreakersCount); i++ {
		var windowBreakerInfo core.WindowBreakerInfo
		if err := windowBreakerInfo.Deserialize(buffer); err != nil {
			return err
		}
		ws.WindowBreakerInfos[i] = windowBreakerInfo
	}
	return nil
}

func (ws *WorldSpawn) SerializeWindowBreakers(buffer *bytes.Buffer) error {
	windowBreakersCount := int32(len(ws.WindowBreakerInfos))
	if err := helpers.WriteInt32(buffer, windowBreakersCount); err != nil {
		return err
	}
	for _, windowBreakerInfo := range ws.WindowBreakerInfos {
		if err := windowBreakerInfo.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ws *WorldSpawn) SynchronizableObjectType(buffer *bytes.Buffer) error {
	var synchronizableObjectTypesCount uint16
	err := helpers.ReadUInt16(buffer, &synchronizableObjectTypesCount)
	if err != nil {
		return err
	}

	ws.SynchronizableObjectTypes = make([]SynchronizableObjectType, 0, synchronizableObjectTypesCount)
	for i := 0; i < int(synchronizableObjectTypesCount); i++ {
		var objectType byte
		if err = helpers.ReadByte(buffer, &objectType); err != nil {
			return err
		}

		switch enums.SynchronizableObjectType(objectType) {
		case enums.AirDrop:
			var airDrop core.AirDrop
			if err := airDrop.Deserialize(buffer); err != nil {
				return err
			}
			ws.SynchronizableObjectTypes = append(ws.SynchronizableObjectTypes, SynchronizableObjectType{Type: enums.SynchronizableObjectType(objectType), Value: &airDrop})
		case enums.AirPlane:
			var airPlane core.AirPlane
			if err := airPlane.Deserialize(buffer); err != nil {
				return err
			}
			ws.SynchronizableObjectTypes = append(ws.SynchronizableObjectTypes, SynchronizableObjectType{Type: enums.SynchronizableObjectType(objectType), Value: &airPlane})
		}
	}
	return nil
}

func (ws *WorldSpawn) SerializeSynchronizableObjectType(buffer *bytes.Buffer) error {
	synchronizableObjectTypesCount := uint16(len(ws.SynchronizableObjectTypes))
	if err := helpers.WriteUInt16(buffer, synchronizableObjectTypesCount); err != nil {
		return err
	}
	for _, item := range ws.SynchronizableObjectTypes {
		if err := helpers.WriteUInt16(buffer, uint16(item.Type)); err != nil {
			return err
		}
		if err := item.Value.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (ws *WorldSpawn) ReadBTR(buffer *bytes.Buffer) error {
	return ws.BTR.Deserialize(buffer)
}

func (ws *WorldSpawn) SerializeBTR(buffer *bytes.Buffer) error {
	return ws.BTR.Serialize(buffer)
}

//ExfiltrationController:  {[{EXFIL_ZB013 2 0 []} {Dorms V-Ex 1 0 []} {ZB-1011 4 0 []} {Crossroads 4 0 []} {Old Gas Station 1 0 []} {Trailer Park 4 0 []} {RUAF Roadblock 4 0 []} {Smuggler's Boat 1 0 []} {ZB-1012 1 0 []}]}
//BufferZoneControllerClass:  {false}
//SmokeGrenades:  []
//Doors:  []
//LampControllers:  [{-1155525479 2} {410558462 2} {7273935 2} {1573357876 2} {1170073349 2} {-1558810006 2} {-1962094533 2} {-396010592 2} {-1511755839 2} {54328102 2} {-1155525480 2} {410558461 2} {7273934 2} {1170073348 2} {-1558810007 2} {-1962094534 2} {-396010593 2} {-1511755840 2} {54328101 2} {-1155525481 2} {7273933 2} {1573357874 2} {1170073347 2} {-1558810008 2} {-1962094535 2} {-396010594 2} {-1511755841 2} {54328100 2} {-1155525482 2} {410558459 2} {7273932 2} {-1962094536 2} {-396010595 2} {-1511755842 2} {54328099 2} {-1155525475 2} {410558466 2} {7273939 2} {1573357880 2} {1170073353 2} {-1962094529 2} {-396010588 2} {54328106 2} {-1155525476 2} {410558465 3} {7273938 3} {1573357879 2} {1170073352 2} {-1558810003 2} {-1962094530 3} {-396010589 2} {-1511755836 2} {54328105 2} {-1155525477 2} {410558464 2} {7273937 2} {1573357878 2} {1170073351 2} {-1558810004 2} {-1962094531 2} {-396010590 2} {-1511755837 2} {54328104 2} {-1155525478 2} {410558463 2} {7273936 2} {1573357877 2} {1170073350 2} {-1558810005 2} {-1962094532 2} {-396010591 2} {-1511755838 2} {54328103 2} {-1155525487 2} {410558454 2} {7273927 2} {1573357868 2} {1170073341 2} {-1558810014 2} {-1962094541 2} {-396010600 2} {-1511755847 2} {54328094 2} {-1155525488 2} {7273926 4} {1573357867 2} {1170073340 2} {-1558810015 4} {-1962094542 2} {-396010601 2} {-1511755848 2} {54328093 2} {1835479902 2} {269395961 2} {-1296687980 2} {1432195375 2} {-133888566 2} {-1699972507 2} {1028910848 2} {-537173093 2} {1479249542 2} {-86834399 2} {1835479901 2} {269395960 2} {-1296687981 2} {1432195374 2} {-133888567 2} {-1699972508 2} {1028910847 2} {-537173094 2} {1479249541 2} {-86834400 2} {1835479900 2} {269395959 2} {-1296687982 2} {1432195373 2} {-133888568 2} {-1699972509 2} {1028910846 2} {-537173095 2} {1479249540 2} {-86834401 2} {269395958 2} {-1296687983 2} {1432195372 2} {-133888569 2} {-1699972510 2} {1028910845 2} {-537173096 2} {1479249539 2} {-86834402 2} {1835479906 2} {269395965 2} {-1296687976 2} {1432195379 2} {-133888562 2} {-1699972503 2} {1028910852 2} {-537173089 2} {1479249546 2} {-86834395 2} {116166903 3} {-1449917038 3} {-1046632511 3} {1682250844 3} {922735957 3} {519451430 3} {2085535371 3} {-1093686678 3} {472397263 3} {-643347984 3} {116166904 3} {-1449917037 3} {-1046632510 3} {1682250845 3} {2085535372 3} {519451431 3} {922735958 3} {-643347983 3} {472397264 3} {-1093686677 3} {116166901 3} {-1449917040 3} {-1046632513 3} {1682250842 3} {2085535369 3} {519451428 3} {922735955 3} {-643347986 3} {472397261 3} {-1093686680 3} {116166902 3} {-1449917039 3} {-1046632512 3} {1682250843 3} {-370620410 2} {-370620411 0}]
//WindowBreakers:  []
//SynchronizableObjectType:  []
//BTR:  {false []}

func (ws *WorldSpawn) GetDefault() error {
	ws.ExfiltrationController = core.ExfiltrationController{
		ExfilDatas: []core.ExfilData{
			{
				Name:               "EXFIL_ZB013",
				ExfiltrationStatus: core.UncompleteRequirements,
			},
			{
				Name:               "Dorms V-Ex",
				ExfiltrationStatus: core.NotPresent,
			},
			{
				Name:               "ZB-1011",
				ExfiltrationStatus: core.RegularMode,
			},
			{
				Name:               "Crossroads",
				ExfiltrationStatus: core.RegularMode,
			},
			{
				Name:               "Old Gas Station",
				ExfiltrationStatus: core.NotPresent,
			},
			{
				Name:               "Trailer Park",
				ExfiltrationStatus: core.RegularMode,
			},
			{
				Name:               "RUAF Roadblock",
				ExfiltrationStatus: core.RegularMode,
			},
			{
				Name:               "Smuggler's Boat",
				ExfiltrationStatus: core.NotPresent,
			},
			{
				Name:               "ZB-1012",
				ExfiltrationStatus: core.NotPresent,
			},
		},
	}

	ws.SmokeGrenades = []core.SmokeGrenadeInfo{}
	ws.DoorInfos = []core.DoorInfo{}

	ws.LampControllerInfos = []core.LampControllerInfo{
		{Id: -1155525479, State: 2},
		{Id: 410558462, State: 2},
		{Id: 7273935, State: 2},
		{Id: 1573357876, State: 2},
		{Id: 1170073349, State: 2},
		{Id: -1558810006, State: 2},
		{Id: -1962094533, State: 2},
		{Id: -396010592, State: 2},
		{Id: -1511755839, State: 2},
		{Id: 54328102, State: 2},
		{Id: -1155525480, State: 2},
		{Id: 410558461, State: 2},
		{Id: 7273934, State: 2},
		{Id: 1170073348, State: 2},
		{Id: -1558810007, State: 2},
		{Id: -1962094534, State: 2},
		{Id: -396010593, State: 2},
		{Id: -1511755840, State: 2},
		{Id: 54328101, State: 2},
		{Id: -1155525481, State: 2},
		{Id: 7273933, State: 2},
		{Id: 1573357874, State: 2},
		{Id: 1170073347, State: 2},
		{Id: -1558810008, State: 2},
		{Id: -1962094535, State: 2},
		{Id: -396010594, State: 2},
		{Id: -1511755841, State: 2},
		{Id: 54328100, State: 2},
		{Id: -1155525482, State: 2},
		{Id: 410558459, State: 2},
		{Id: 7273932, State: 2},
		{Id: -1962094536, State: 2},
		{Id: -396010595, State: 2},
		{Id: -1511755842, State: 2},
		{Id: 54328099, State: 2},
		{Id: -1155525475, State: 2},
		{Id: 410558466, State: 2},
		{Id: 7273939, State: 2},
		{Id: 1573357880, State: 2},
		{Id: 1170073353, State: 2},
		{Id: -1962094529, State: 2},
		{Id: -396010588, State: 2},
		{Id: 54328106, State: 2},
		{Id: -1155525476, State: 2},
		{Id: 410558465, State: 3},
		{Id: 7273938, State: 3},
		{Id: 1573357879, State: 2},
		{Id: 1170073352, State: 2},
		{Id: -1558810003, State: 2},
		{Id: -1962094530, State: 3},
		{Id: -396010589, State: 2},
		{Id: -1511755836, State: 2},
		{Id: 54328105, State: 2},
		{Id: -1155525477, State: 2},
		{Id: 410558464, State: 2},
		{Id: 7273937, State: 2},
		{Id: 1573357878, State: 2},
		{Id: 1170073351, State: 2},
		{Id: -1558810004, State: 2},
		{Id: -1962094531, State: 2},
		{Id: -396010590, State: 2},
		{Id: -1511755837, State: 2},
		{Id: 54328104, State: 2},
		{Id: -1155525478, State: 2},
		{Id: 410558463, State: 2},
		{Id: 7273936, State: 2},
		{Id: 1573357877, State: 2},
		{Id: 1170073350, State: 2},
		{Id: -1558810005, State: 2},
		{Id: -1962094532, State: 2},
		{Id: -396010591, State: 2},
		{Id: -1511755838, State: 2},
		{Id: 54328103, State: 2},
		{Id: -1155525487, State: 2},
		{Id: 410558454, State: 2},
		{Id: 7273927, State: 2},
		{Id: 1573357868, State: 2},
		{Id: 1170073341, State: 2},
		{Id: -1558810014, State: 2},
		{Id: -1962094541, State: 2},
		{Id: -396010600, State: 2},
		{Id: -1511755847, State: 2},
		{Id: 54328094, State: 2},
		{Id: -1155525488, State: 2},
		{Id: 7273926, State: 4},
		{Id: 1573357867, State: 2},
		{Id: 1170073340, State: 2},
		{Id: -1558810015, State: 4},
		{Id: -1962094542, State: 2},
		{Id: -396010601, State: 2},
		{Id: -1511755848, State: 2},
		{Id: 54328093, State: 2},
		{Id: 1835479902, State: 2},
		{Id: 269395961, State: 2},
		{Id: -1296687980, State: 2},
		{Id: 1432195375, State: 2},
		{Id: -133888566, State: 2},
		{Id: -1699972507, State: 2},
		{Id: 1028910848, State: 2},
		{Id: -537173093, State: 2},
		{Id: 1479249542, State: 2},
		{Id: -86834399, State: 2},
		{Id: 1835479901, State: 2},
		{Id: 269395960, State: 2},
		{Id: -1296687981, State: 2},
		{Id: 1432195374, State: 2},
		{Id: -133888567, State: 2},
		{Id: -1699972508, State: 2},
		{Id: 1028910847, State: 2},
		{Id: -537173094, State: 2},
		{Id: 1479249541, State: 2},
		{Id: -86834400, State: 2},
		{Id: 1835479900, State: 2},
		{Id: 269395959, State: 2},
		{Id: -1296687982, State: 2},
		{Id: 1432195373, State: 2},
		{Id: -133888568, State: 2},
		{Id: -1699972509, State: 2},
		{Id: 1028910846, State: 2},
		{Id: -537173095, State: 2},
		{Id: 1479249540, State: 2},
		{Id: -86834401, State: 2},
		{Id: 269395958, State: 2},
		{Id: -1296687983, State: 2},
		{Id: 1432195372, State: 2},
		{Id: -133888569, State: 2},
		{Id: -1699972510, State: 2},
		{Id: 1028910845, State: 2},
		{Id: -537173096, State: 2},
		{Id: 1479249539, State: 2},
		{Id: -86834402, State: 2},
		{Id: 1835479906, State: 2},
		{Id: 269395965, State: 2},
		{Id: -1296687976, State: 2},
		{Id: 1432195379, State: 2},
		{Id: -133888562, State: 2},
		{Id: -1699972503, State: 2},
		{Id: 1028910852, State: 2},
		{Id: -537173089, State: 2},
		{Id: 1479249546, State: 2},
		{Id: -86834395, State: 2},
		{Id: 116166903, State: 3},
		{Id: -1449917038, State: 3},
		{Id: -1046632511, State: 3},
		{Id: 1682250844, State: 3},
		{Id: 922735957, State: 3},
		{Id: 519451430, State: 3},
		{Id: 2085535371, State: 3},
		{Id: -1093686678, State: 3},
		{Id: 472397263, State: 3},
		{Id: -643347984, State: 3},
		{Id: 116166904, State: 3},
		{Id: -1449917037, State: 3},
		{Id: -1046632510, State: 3},
		{Id: 1682250845, State: 3},
		{Id: 2085535372, State: 3},
		{Id: 519451431, State: 3},
		{Id: 922735958, State: 3},
		{Id: -643347983, State: 3},
		{Id: 472397264, State: 3},
		{Id: -1093686677, State: 3},
		{Id: 116166901, State: 3},
		{Id: -1449917040, State: 3},
		{Id: -1046632513, State: 3},
		{Id: 1682250842, State: 3},
		{Id: 2085535369, State: 3},
		{Id: 519451428, State: 3},
		{Id: 922735955, State: 3},
		{Id: -643347986, State: 3},
		{Id: 472397261, State: 3},
		{Id: -1093686680, State: 3},
		{Id: 116166902, State: 3},
		{Id: -1449917039, State: 3},
		{Id: -1046632512, State: 3},
		{Id: 1682250843, State: 3},
		{Id: -370620410, State: 2},
		{Id: -370620411, State: 0},
	}
	ws.WindowBreakerInfos = []core.WindowBreakerInfo{}
	ws.SynchronizableObjectTypes = []SynchronizableObjectType{}
	ws.BTR = core.BTR{
		HasBTR:  false,
		BTRData: []byte{},
	}
	return nil
}
