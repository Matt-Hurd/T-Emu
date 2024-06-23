package core

import (
	"bytes"
	"game-server/helpers"
)

type SmokeGrenadeInfo struct {
	Id          string
	Position    Vector3
	Template    string
	Time        int32
	Orientation Quaternion
	PlatformId  int16
}

func (s *SmokeGrenadeInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadString(buffer, &s.Id); err != nil {
		return err
	}
	if err = s.Position.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &s.Template); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &s.Time); err != nil {
		return err
	}
	if err = s.Orientation.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadInt16(buffer, &s.PlatformId); err != nil {
		return err
	}
	return nil
}

func (s *SmokeGrenadeInfo) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteString(buffer, s.Id); err != nil {
		return err
	}
	if err = s.Position.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteString(buffer, s.Template); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, s.Time); err != nil {
		return err
	}
	if err = s.Orientation.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteInt16(buffer, s.PlatformId); err != nil {
		return err
	}
	return nil
}
