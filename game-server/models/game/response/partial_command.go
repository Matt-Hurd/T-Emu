package response

import (
	"bytes"
	"game-server/helpers"
)

type PacketPartialCommand struct {
	CommandKey int16
	Id         int32
	PartNum    byte
	Offset     int32
	PartsCount byte
	PartSize   uint16
	Size       int32
	Data       []byte
}

func (p *PacketPartialCommand) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadInt16(buffer, &p.CommandKey); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.Id); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &p.PartNum); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.Offset); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &p.PartsCount); err != nil {
		return err
	}
	if err = helpers.ReadUInt16(buffer, &p.PartSize); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.Size); err != nil {
		return err
	}
	p.Data = make([]byte, p.PartSize)
	if _, err = buffer.Read(p.Data); err != nil {
		return err
	}
	return nil
}

func (p *PacketPartialCommand) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteInt16(buffer, p.CommandKey); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, p.Id); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, p.PartNum); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, p.Offset); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, p.PartsCount); err != nil {
		return err
	}
	if err = helpers.WriteUInt16(buffer, p.PartSize); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, p.Size); err != nil {
		return err
	}
	if _, err = buffer.Write(p.Data); err != nil {
		return err
	}
	return nil
}
