package response

import (
	"bytes"
	"game-server/helpers"
)

type PacketSpawnFinished struct {
	State uint32
}

func (p *PacketSpawnFinished) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadPackedUInt32(buffer, &p.State)
	if err != nil {
		panic(err)
	}
	return nil
}

func (p *PacketSpawnFinished) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WritePackedUInt32(buffer, uint32(p.State))
	if err != nil {
		return err
	}
	return nil
}
