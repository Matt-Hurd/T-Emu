package math

import (
	"bytes"
	"game-server/helpers"

	"github.com/g3n/engine/math32"
)

type Vector3 struct {
	math32.Vector3
}

func (v *Vector3) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadFloat32(buffer, &v.X); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &v.Y); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &v.Z); err != nil {
		return err
	}
	return nil
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

func (v *Vector3) ToQuaternion() *Quaternion {
	return &Quaternion{math32.Quaternion{X: v.X, Y: v.Y, Z: v.Z, W: 0}}
}
