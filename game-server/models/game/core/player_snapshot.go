package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
	"game-server/models/game/math"

	"github.com/g3n/engine/math32"
)

type PlayerSnapshot struct {
	ClientTime           float32
	BodyPosition         math.Vector3
	BodyRotation         math.Vector2
	HeadRotation         math.Vector2
	MovementDirection    math.Vector2
	Velocity             math.Vector3
	Tilt                 float32
	Step                 int32
	BlindFire            int32
	State                enums.PlayerState
	StateAnimatorIndex   byte
	PhysicalCondition    enums.PhysicalCondition
	MovementSpeed        float32
	SprintSpeed          float32
	MaxSpeed             float32
	Pose                 enums.PlayerPose
	PoseLevel            float32
	InHandsObjectOverlap float32
	mask                 byte
	AttachedToMask       bool
	JumpHeight           float32
	FallHeight           float32
	FallTime             float32
	LeftStance           bool
}

func (p *PlayerSnapshot) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadFloat32(buffer, &p.ClientTime); err != nil {
		return err
	}
	if err = p.BodyPosition.Deserialize(buffer); err != nil {
		return err
	}
	x, y := float32(0), int16(0)
	if err = helpers.ReadFloat32(buffer, &x); err != nil {
		return err
	}
	if err = helpers.ReadInt16(buffer, &y); err != nil {
		return err
	}
	p.BodyRotation = math.Vector2{
		Vector2: math32.Vector2{X: x, Y: math.ScaleShortToFloat(y, -90, -90)},
	}
	var headX, headY byte
	if err = helpers.ReadByte(buffer, &headX); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &headY); err != nil {
		return err
	}
	p.HeadRotation = math.Vector2{
		Vector2: math32.Vector2{X: math.ScaleByteToFloat(headX, -50, 20), Y: math.ScaleByteToFloat(headY, -40, 40)},
	}
	var moveX, moveY byte
	if err = helpers.ReadByte(buffer, &moveX); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &moveY); err != nil {
		return err
	}
	p.MovementDirection = math.ScaleFromVector2Byte([]byte{moveX, moveY}, -1, 1)
	var velX, velY, velZ int16
	if err = helpers.ReadInt16(buffer, &velX); err != nil {
		return err
	}
	if err = helpers.ReadInt16(buffer, &velY); err != nil {
		return err
	}
	if err = helpers.ReadInt16(buffer, &velZ); err != nil {
		return err
	}
	p.Velocity = math.ScaleFromVector3Short([]int16{velX, velY, velZ}, -25, 25)
	var tilt byte
	if err = helpers.ReadByte(buffer, &tilt); err != nil {
		return err
	}
	p.Tilt = math.ScaleByteToFloat(tilt, -5, 5)
	var step byte
	if err = helpers.ReadByte(buffer, &step); err != nil {
		return err
	}
	p.Step = math.ScaleByteToInt(step, -1, 1)
	var blindFire byte
	if err = helpers.ReadByte(buffer, &blindFire); err != nil {
		return err
	}
	p.BlindFire = math.ScaleByteToInt(blindFire, -1, 1)
	if err = helpers.ReadByte(buffer, (*byte)(&p.State)); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &p.StateAnimatorIndex); err != nil {
		return err
	}
	var physicalCondition byte
	if err = helpers.ReadByte(buffer, &physicalCondition); err != nil {
		return err
	}
	p.PhysicalCondition = enums.PhysicalCondition(physicalCondition)
	var movementSpeed byte
	if err = helpers.ReadByte(buffer, &movementSpeed); err != nil {
		return err
	}
	p.MovementSpeed = math.ScaleByteToFloat(movementSpeed, 0, 1)
	var sprintSpeed byte
	if err = helpers.ReadByte(buffer, &sprintSpeed); err != nil {
		return err
	}
	p.SprintSpeed = math.ScaleByteToFloat(sprintSpeed, 0, 1)
	var maxSpeed byte
	if err = helpers.ReadByte(buffer, &maxSpeed); err != nil {
		return err
	}
	p.MaxSpeed = math.ScaleByteToFloat(maxSpeed, 0, 1)
	if err = helpers.ReadByte(buffer, (*byte)(&p.Pose)); err != nil {
		return err
	}
	var poseLevel byte
	if err = helpers.ReadByte(buffer, &poseLevel); err != nil {
		return err
	}
	p.PoseLevel = math.ScaleByteToFloat(poseLevel, 0, 1)
	var inHandsObjectOverlap byte
	if err = helpers.ReadByte(buffer, &inHandsObjectOverlap); err != nil {
		return err
	}
	p.InHandsObjectOverlap = math.ScaleByteToFloat(inHandsObjectOverlap, 0, 1)
	if err = helpers.ReadByte(buffer, &p.mask); err != nil {
		return err
	}
	p.AttachedToMask = (p.mask == (p.mask | 1<<0))
	var jumpHeight byte
	if err = helpers.ReadByte(buffer, &jumpHeight); err != nil {
		return err
	}
	p.JumpHeight = math.ScaleByteToFloat(jumpHeight, -10, 10)
	var fallHeight byte
	if err = helpers.ReadByte(buffer, &fallHeight); err != nil {
		return err
	}
	p.FallHeight = math.ScaleByteToFloat(fallHeight, 0, 10)
	var fallTime byte
	if err = helpers.ReadByte(buffer, &fallTime); err != nil {
		return err
	}
	p.FallTime = math.ScaleByteToFloat(fallTime, 0, 10)
	if err = helpers.ReadBool(buffer, &p.LeftStance); err != nil {
		return err
	}
	return nil
}

func (p *PlayerSnapshot) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteFloat32(buffer, p.ClientTime); err != nil {
		return err
	}
	if err = p.BodyPosition.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, p.BodyRotation.X); err != nil {
		return err
	}
	if err = helpers.WriteInt16(buffer, math.ScaleFloatToShort(p.BodyRotation.Y, -90, -90)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.HeadRotation.X, -50, 20)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.HeadRotation.Y, -40, 40)); err != nil {
		return err
	}
	movementDirection := math.ScaleToVector2Byte(p.MovementDirection, -1, 1)
	if err = helpers.WriteByte(buffer, movementDirection[0]); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, movementDirection[1]); err != nil {
		return err
	}
	velocity := math.ScaleToVector3Short(p.Velocity, -25, 25)
	if err = helpers.WriteInt16(buffer, velocity[0]); err != nil {
		return err
	}
	if err = helpers.WriteInt16(buffer, velocity[1]); err != nil {
		return err
	}
	if err = helpers.WriteInt16(buffer, velocity[2]); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.Tilt, -5, 5)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleIntToByte(p.Step, -1, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleIntToByte(p.BlindFire, -1, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(p.State)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, p.StateAnimatorIndex); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(p.PhysicalCondition)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.MovementSpeed, 0, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.SprintSpeed, 0, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.MaxSpeed, 0, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(p.Pose)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.PoseLevel, 0, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.InHandsObjectOverlap, 0, 1)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, p.mask); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.JumpHeight, -10, 10)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.FallHeight, 0, 10)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, math.ScaleFloatToByte(p.FallTime, 0, 10)); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, p.LeftStance); err != nil {
		return err
	}
	return nil
}
