package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type CorpseImpulse struct {
	BodyPartColliderType enums.BodyPartColliderType
	Direction            Vector3
	Point                Vector3
	Force                float32
	OverallVelocity      float32
}

func (impulse *CorpseImpulse) Serialize(buffer *bytes.Buffer) error {
	helpers.WriteInt16(buffer, int16(impulse.BodyPartColliderType))
	err := impulse.Direction.Serialize(buffer)
	if err != nil {
		return err
	}
	err = impulse.Point.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, impulse.Force)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, impulse.OverallVelocity)
	if err != nil {
		return err
	}
	return nil
}

func (impulse *CorpseImpulse) Deserialize(buffer *bytes.Buffer) error {
	tmp := int16(0)
	err := helpers.ReadInt16(buffer, &tmp)
	impulse.BodyPartColliderType = enums.BodyPartColliderType(tmp)
	if err != nil {
		return err
	}
	err = impulse.Direction.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = impulse.Point.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &impulse.Force)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &impulse.OverallVelocity)
	if err != nil {
		return err
	}
	return nil
}
