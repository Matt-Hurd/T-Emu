package core

import (
	"bytes"
	"game-server/helpers"
)

type Vector3 struct {
	X, Y, Z float32
}

func DeserializeVector3(buffer *bytes.Buffer) (Vector3, error) {
	var v Vector3
	var err error

	if err = helpers.ReadFloat32(buffer, &v.X); err != nil {
		return v, err
	}
	if err = helpers.ReadFloat32(buffer, &v.Y); err != nil {
		return v, err
	}
	if err = helpers.ReadFloat32(buffer, &v.Z); err != nil {
		return v, err
	}
	return v, nil
}

func (v *Vector3) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteFloat32(buffer, v.X); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, v.Y); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, v.Z); err != nil {
		return err
	}
	return nil
}
