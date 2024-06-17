package game

import (
	"bytes"
	"game-server/helpers"
)

type PacketConnectionRequest struct {
	ProfileID        string
	Token            string
	ObserveOnly      bool
	EncryptionKey    []byte
	EncryptionKeyLen uint32
	LocationId       string
}

func (p *PacketConnectionRequest) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadString(buffer, &p.ProfileID); err != nil {
		return err
	}

	if err := helpers.ReadString(buffer, &p.Token); err != nil {
		return err
	}

	if err := helpers.ReadBool(buffer, &p.ObserveOnly); err != nil {
		return err
	}

	if err := helpers.ReadBytesAndSize(buffer, &p.EncryptionKey); err != nil {
		return err
	}

	if err := helpers.ReadUInt32(buffer, &p.EncryptionKeyLen); err != nil {
		return err
	}

	if err := helpers.ReadString(buffer, &p.LocationId); err != nil {
		return err
	}

	// fmt.Printf("Parsed PacketConnectionRequest: %v\n", p)

	return nil
}
