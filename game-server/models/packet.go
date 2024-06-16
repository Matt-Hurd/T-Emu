package models

import (
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
	Parse([]byte) error
}

func ParsePacket(data []byte) (Packet, error) {
	if len(data) < 1 {
		return nil, errors.New("invalid packet length")
	}
	packetType := data[0]
	packetData := data[1:]

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

	if err := packet.Parse(packetData); err != nil {
		return nil, err
	}

	return packet, nil
}
