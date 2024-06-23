package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"game-server/helpers"
	"game-server/models/game/request"
	"game-server/models/game/response"
)

type GamePacket interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

type DataPacket struct {
	Length         uint16
	GamePacketType uint16
	GamePacket     GamePacket
}

func (p *DataPacket) Type() byte {
	return PacketTypeData
}

func (p *DataPacket) Parse(buffer *bytes.Buffer) error {
	helpers.ReadUInt16(buffer, &p.Length)
	helpers.ReadUInt16(buffer, &p.GamePacketType)

	packetBuf := make([]byte, p.Length)
	n, err := buffer.Read(packetBuf)

	if err != nil {
		return err
	}

	if int(p.Length) != n {
		fmt.Printf("IGNORING ERROR Length: %d, Buffer Length: %d\n", p.Length, n)
		// fmt.Printf("ERRORING data: %x\n", data)
		// return fmt.Errorf("ERROR Length: %d, Buffer Length: %d", p.Length, buffer.Len())
	}

	tmpBuffer := bytes.NewBuffer(packetBuf)

	var res interface {
		Serialize(buffer *bytes.Buffer) error
		Deserialize(buffer *bytes.Buffer) error
	}
	switch p.GamePacketType {
	case 2:
		res = &response.PacketRpcResponse{}
		// return fmt.Errorf("PacketRpcResponse unexpected")
	case 3:
		res = &response.PacketObjectSpawn{}
	case 5:
		res = &request.PacketCmdRequest{}
	case 12:
		res = &response.PacketSpawnFinished{}
	case 15:
		res = &response.PacketClientAuthority{}
	case 35:
		res = &request.PacketClientReady{}
	case 147:
		res = &request.PacketConnection{}
	case 151:
		res = &response.WorldSpawn{}
	case 174:
		res = &response.PacketCommandsObservedPlayers{}
	case 190:
		res = &request.PacketProgressReport{}
	case 191:
		res = &response.SubWorldSpawnLoot{}
	case 192:
		res = &response.SubWorldSpawnSearchLoot{}
	case 18385:
		res = &request.PacketHLAPIRequest{}
	default:
		return nil
	}

	err = res.Deserialize(tmpBuffer)
	if err != nil {
		return err
	}
	if p.GamePacketType != 5 && p.GamePacketType != 3 {
		fmt.Printf("Received data packet type: %d. %v\n", p.GamePacketType, res)
	}
	p.GamePacket = res
	return nil
}

func (p *DataPacket) Write() []byte {
	buffer := bytes.Buffer{}
	err := p.GamePacket.Serialize(&buffer)
	if err != nil {
		fmt.Println("Error serializing data packet:", err)
	}
	p.Length = uint16(buffer.Len())

	fmt.Printf("Writing data packet type: %d\n", p.GamePacketType)

	data := make([]byte, 4)
	binary.LittleEndian.PutUint16(data[:2], uint16(buffer.Len()))
	binary.LittleEndian.PutUint16(data[2:], uint16(p.GamePacketType))
	return append(data, buffer.Bytes()...)
}
