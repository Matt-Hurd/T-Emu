package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"game-server/helpers"
	"game-server/models/game/request"
)

type DataPacket struct {
	Length         uint16
	GamePacketType uint16
	GamePacket     interface {
		Serialize(buffer *bytes.Buffer) error
		Deserialize(buffer *bytes.Buffer) error
	}
}

func (p *DataPacket) Type() byte {
	return PacketTypeData
}

func (p *DataPacket) Parse(data []byte) error {
	buffer := bytes.NewBuffer(data)
	helpers.ReadUInt16(buffer, &p.Length)
	helpers.ReadUInt16(buffer, &p.GamePacketType)

	if int(p.Length) != buffer.Len() {
		fmt.Printf("ERROR Length: %d, Buffer Length: %d\n", p.Length, buffer.Len())
		return fmt.Errorf("ERROR Length: %d, Buffer Length: %d", p.Length, buffer.Len())
	}

	var res interface {
		Serialize(buffer *bytes.Buffer) error
		Deserialize(buffer *bytes.Buffer) error
	}
	switch p.GamePacketType {
	case 2:
		// res = &response.PacketRpcResponse{}
		return fmt.Errorf("PacketRpcResponse unexpected")
	case 5:
		res = &request.PacketCmdRequest{}
	case 147:
		res = &request.PacketConnection{}
	case 190:
		res = &request.PacketProgressReport{}
	default:
		return nil
	}

	err := res.Deserialize(buffer)
	if err != nil {
		return err
	}
	p.GamePacket = res
	return nil
}

func (p *DataPacket) Write() []byte {
	buffer := bytes.Buffer{}
	err := p.GamePacket.Serialize(&buffer)
	if err != nil {
		fmt.Println("Error serializing connection response:", err)
	}
	p.Length = uint16(buffer.Len())

	data := make([]byte, 4)
	binary.LittleEndian.PutUint16(data[:2], uint16(buffer.Len()))
	binary.LittleEndian.PutUint16(data[2:], uint16(p.GamePacketType))
	return append(data, buffer.Bytes()...)
}
