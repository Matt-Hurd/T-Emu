package request

import (
	"bytes"
	"game-server/helpers"
)

type PacketProgressReport struct {
	ProfileID string
	Id        int32
	Progress  float32
}

func (p *PacketProgressReport) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadString(buffer, &p.ProfileID); err != nil {
		return err
	}

	if err := helpers.ReadInt32(buffer, &p.Id); err != nil {
		return err
	}

	if err := helpers.ReadFloat32(buffer, &p.Progress); err != nil {
		return err
	}
	return nil
}

func (p *PacketProgressReport) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, p.ProfileID); err != nil {
		return err
	}

	if err := helpers.WriteInt32(buffer, p.Id); err != nil {
		return err
	}

	if err := helpers.WriteFloat32(buffer, p.Progress); err != nil {
		return err
	}

	return nil
}
