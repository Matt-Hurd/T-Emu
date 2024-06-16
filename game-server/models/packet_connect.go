package models

import "bytes"

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

func (p *ConnectPacket) Write() []byte {
	output := make([]byte, 2)
	output[0] = p.Type()
	output[1] = 1
	return output
}
