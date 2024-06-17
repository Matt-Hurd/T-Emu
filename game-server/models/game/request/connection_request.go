package request

import (
	"bytes"
	"game-server/helpers"
)

type PacketConnection struct {
	ProfileID        string
	Token            string
	ObserveOnly      bool
	EncryptionKey    []byte
	EncryptionKeyLen uint32
	LocationId       string
}

func (p *PacketConnection) Deserialize(buffer *bytes.Buffer) error {
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

	// fmt.Printf("Parsed PacketConnection: %v\n", p)

	return nil
}

func (p *PacketConnection) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, p.ProfileID); err != nil {
		return err
	}

	if err := helpers.WriteString(buffer, p.Token); err != nil {
		return err
	}

	if err := helpers.WriteBool(buffer, p.ObserveOnly); err != nil {
		return err
	}

	if err := helpers.WriteBytesAndSize(buffer, p.EncryptionKey); err != nil {
		return err
	}

	if err := helpers.WriteUInt32(buffer, p.EncryptionKeyLen); err != nil {
		return err
	}

	if err := helpers.WriteString(buffer, p.LocationId); err != nil {
		return err
	}

	return nil
}
