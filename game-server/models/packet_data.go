package models

import (
	"bytes"
	"fmt"
	"game-server/models/game"
)

type DataPacket struct {
	Length         int16
	GamePacketType int16
	GamePacket     interface{}
}

func (p *DataPacket) Type() byte {
	return PacketTypeData
}

func (p *DataPacket) Parse(data []byte) error {
	buffer := bytes.NewBuffer(data)
	buffer.Read(data[:2])
	p.Length = int16(data[0]) | int16(data[1])<<8

	if int(p.Length) != buffer.Len()-2 {
		fmt.Printf("ERROR Length: %d, Buffer Length: %d\n", p.Length, buffer.Len())
		return fmt.Errorf("ERROR Length: %d, Buffer Length: %d", p.Length, buffer.Len())
	}

	buffer.Read(data[:2])
	msgtype := int16(data[0]) | int16(data[1])<<8
	if msgtype != 147 {
		return fmt.Errorf("ERROR Message Type: %d", msgtype)
	}
	p.GamePacketType = msgtype

	switch p.GamePacketType {
	case 147:
		res := &game.PacketConnectionRequest{}
		err := res.Parse(buffer)
		if err != nil {
			return err
		}
		p.GamePacket = res
		return nil
	default:
		fmt.Printf("Unknown GamePacketType: %d\n", p.GamePacketType)
		return fmt.Errorf("unknown GamePacketType: %d", p.GamePacketType)
	}
}

func (p *DataPacket) Write() []byte {
	return nil
}
