package math

import (
	"bytes"
	"game-server/helpers"

	"github.com/g3n/engine/math32"
)

type Quaternion struct {
	math32.Quaternion
}

func (q *Quaternion) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadFloat32(buffer, &q.X); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &q.Y); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &q.Z); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &q.W); err != nil {
		return err
	}
	return nil
}

func (q *Quaternion) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteFloat32(buffer, q.X); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, q.Y); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, q.Z); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, q.W); err != nil {
		return err
	}
	return nil
}

func (q *Quaternion) ToVector3() *Vector3 {
	return &Vector3{math32.Vector3{X: q.X, Y: q.Y, Z: q.Z}}
}
