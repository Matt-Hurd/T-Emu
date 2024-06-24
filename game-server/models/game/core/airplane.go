package core

import (
	"bytes"
	"game-server/helpers"
	"math"

	gameMath "game-server/models/game/math"

	"github.com/g3n/engine/math32"
)

type AirPlane struct {
	Id       uint16
	Position gameMath.Vector3
	Rotation gameMath.Quaternion
	UniqueId int
}

func (a *AirPlane) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadUInt16(buffer, &a.Id); err != nil {
		return err
	}
	if err = a.Position.Deserialize(buffer); err != nil {
		return err
	}
	var rotationVec gameMath.Vector3
	if err = rotationVec.Deserialize(buffer); err != nil {
		return err
	}
	a.Rotation = airplaneToQuaternion(rotationVec)
	var airdropType byte
	if err = helpers.ReadByte(buffer, &airdropType); err != nil {
		return err
	} else {
		a.UniqueId = int(airdropType)
	}
	return nil
}

func (a *AirPlane) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteUInt16(buffer, a.Id); err != nil {
		return err
	}
	if err = a.Position.Serialize(buffer); err != nil {
		return err
	}
	if err = airplaneToVector3(&a.Rotation).Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(a.UniqueId)); err != nil {
		return err
	}
	return nil
}

func airplaneToQuaternion(v gameMath.Vector3) gameMath.Quaternion {
	cy := float32(math.Cos(float64(v.Z) * 0.5))
	sy := float32(math.Sin(float64(v.Z) * 0.5))
	cp := float32(math.Cos(float64(v.Y) * 0.5))
	sp := float32(math.Sin(float64(v.Y) * 0.5))
	cr := float32(math.Cos(float64(v.X) * 0.5))
	sr := float32(math.Sin(float64(v.X) * 0.5))

	return gameMath.Quaternion{
		Quaternion: math32.Quaternion{
			W: cr*cp*cy + sr*sp*sy,
			X: sr*cp*cy - cr*sp*sy,
			Y: cr*sp*cy + sr*cp*sy,
			Z: cr*cp*sy - sr*sp*cy,
		},
	}
}

func airplaneToVector3(q *gameMath.Quaternion) *gameMath.Vector3 {
	// Convert quaternion back to Euler angles (roll, pitch, yaw)
	// This assumes the quaternion is normalized
	ysqr := q.Y * q.Y

	t0 := 2.0 * (q.W*q.X + q.Y*q.Z)
	t1 := 1.0 - 2.0*(q.X*q.X+ysqr)
	roll := float32(math.Atan2(float64(t0), float64(t1)))

	t2 := 2.0 * (q.W*q.Y - q.Z*q.X)
	t2 = float32(math.Max(float64(t2), -1.0))
	t2 = float32(math.Min(float64(t2), 1.0))
	pitch := float32(math.Asin(float64(t2)))

	t3 := 2.0 * (q.W*q.Z + q.X*q.Y)
	t4 := 1.0 - 2.0*(ysqr+q.Z*q.Z)
	yaw := float32(math.Atan2(float64(t3), float64(t4)))

	return &gameMath.Vector3{
		Vector3: math32.Vector3{
			X: roll,
			Y: pitch,
			Z: yaw,
		},
	}
}
