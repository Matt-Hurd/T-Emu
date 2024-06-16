package models

import "game-server/helpers"

type PongPacket struct {
	ElapsedMilliseconds uint32
}

func (p *PongPacket) Type() byte {
	return PacketTypePong
}

func (p *PongPacket) Parse(data []byte) error {
	return nil
}

func (p *PongPacket) Write() []byte {
	return append([]byte{p.Type()}, helpers.UInt32ToBytes(p.ElapsedMilliseconds)...)

}
