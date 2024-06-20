package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PacketObjectSpawn struct {
	NetId    uint32
	AssetId  core.NetworkHash128
	Position core.Vector3
	Payload  []byte
	Rotation core.Quaternion
}

func (p *PacketObjectSpawn) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadPackedUInt32(buffer, &p.NetId)
	if err != nil {
		return err
	}
	err = p.AssetId.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = p.Position.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.ReadBytesAndSize(buffer, &p.Payload)
	if err != nil {
		return err
	}
	err = p.Rotation.Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (p *PacketObjectSpawn) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WritePackedUInt32(buffer, p.NetId)
	if err != nil {
		return err
	}
	err = p.AssetId.Serialize(buffer)
	if err != nil {
		return err
	}
	err = p.Position.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteBytesAndSize(buffer, p.Payload)
	if err != nil {
		return err
	}
	err = p.Rotation.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}
