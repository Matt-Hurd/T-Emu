package response

import (
	"bytes"
	"fmt"
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
	fmt.Println("ExfiltrationController: ", ws.ExfiltrationController)

	if err := ws.BufferZoneControllerClass.Deserialize(buffer); err != nil {
		return err
	}
	fmt.Println("BufferZoneControllerClass: ", ws.BufferZoneControllerClass)

	if err := ws.InitializeSmokeGrenades(buffer); err != nil {
		return err
	}
	fmt.Println("SmokeGrenades: ", ws.SmokeGrenades)

	if err := ws.InitializeDoors(buffer); err != nil {
		return err
	}
	fmt.Println("Doors: ", ws.DoorInfos)

	if err := ws.InitializeLampControllers(buffer); err != nil {
		return err
	}
	fmt.Println("LampControllers: ", ws.LampControllerInfos)

	if err := ws.InitializeWindowBreakers(buffer); err != nil {
		return err
	}
	fmt.Println("WindowBreakers: ", ws.WindowBreakerInfos)

	if err := ws.SynchronizableObjectType(buffer); err != nil {
		return err
	}
	fmt.Println("SynchronizableObjectType: ", ws.SynchronizableObjectTypes)

	if err := ws.ReadBTR(buffer); err != nil {
		return err
	}
	fmt.Println("BTR: ", ws.BTR)

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
