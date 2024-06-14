package models

import (
	"bytes"
	"errors"
)

// Define various packet types
const (
	PacketTypeConnect = iota + 1
)

// Packet interface
type Packet interface {
	Type() byte
	Parse([]byte) error
}

// LoginPacket structure
type ConnectPacket struct {
	Syn byte
	Asc byte
}

func (p *ConnectPacket) Type() byte {
	return PacketTypeConnect
}

func (p *ConnectPacket) Parse(data []byte) error {
	buffer := bytes.NewBuffer(data)
	syn, err := buffer.ReadByte()
	if err != nil {
		return err
	}
	asc, err2 := buffer.ReadByte()
	if err2 != nil {
		return err
	}
	p.Syn = syn
	p.Asc = asc
	return nil
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
	default:
		return nil, errors.New("unknown packet type")
	}

	if err := packet.Parse(packetData); err != nil {
		return nil, err
	}

	return packet, nil
}
