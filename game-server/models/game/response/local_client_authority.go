package response

import (
	"bytes"
	"game-server/helpers"
)

type PacketClientAuthority struct {
	NetId     uint32
	Authority bool
}

func (p *PacketClientAuthority) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadPackedUInt32(buffer, &p.NetId)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &p.Authority)
	if err != nil {
		return err
	}
	return nil
}

func (p *PacketClientAuthority) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WritePackedUInt32(buffer, p.NetId)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, p.Authority)
	if err != nil {
		return err
	}
	return nil
}
