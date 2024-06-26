package models

import (
	"bytes"
	"errors"
	"fmt"
)

// Define various packet types
const (
	PacketTypeConnect = iota + 1
	PacketTypePing
	PacketTypePong
	PacketTypeData
	PacketTypeDisconnect
	PacketTypeFragment = iota + 0x80
	PacketTypeFragmentEnd
)

// Packet interface
type Packet interface {
	Type() byte
	Parse(buffer *bytes.Buffer) error
}

func ParsePacket(buffer *bytes.Buffer) (Packet, error) {
	if buffer.Len() < 1 {
		return nil, errors.New("invalid packet length")
	}
	packetType, err := buffer.ReadByte()
	if err != nil {
		return nil, err
	}

	var packet Packet
	switch packetType {
	case PacketTypeConnect:
		packet = &ConnectPacket{}
	case PacketTypePing:
		packet = &PingPacket{}
	case PacketTypeData:
		packet = &DataPacket{}
	default:
		fmt.Printf("Unknown packet type: %d\n", packetType)
		return nil, errors.New("unknown packet type")
	}

	if err := packet.Parse(buffer); err != nil {
		return nil, err
	}

	return packet, nil
}
