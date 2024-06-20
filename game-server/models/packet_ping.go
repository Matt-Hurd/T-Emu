package models

import (
	"bytes"
	"fmt"
	"game-server/helpers"
)

type PingPacket struct {
	ElapsedMilliseconds uint32
}

func (p *PingPacket) Type() byte {
	return PacketTypePing
}

func (p *PingPacket) Parse(buffer *bytes.Buffer) error {
	err := helpers.ReadUInt32(buffer, &p.ElapsedMilliseconds)
	if err != nil {
		return err
	}

	fmt.Printf("Parsed ElapsedMilliseconds (LittleEndian): %d\n", p.ElapsedMilliseconds)

	return nil
}

func (p *PingPacket) Write(milliseconds uint32) []byte {
	return append([]byte{p.Type()}, helpers.UInt32ToBytes(milliseconds)...)
}
