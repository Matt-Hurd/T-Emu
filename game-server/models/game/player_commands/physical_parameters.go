package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandPhysicalParameters struct {
	Overweight       float32
	TransitionSpeed  float32
	IsHeavyBreathing bool
	WalkOverweight   float32
	SoundRadius      float32
	BreathIsAudible  bool
	MinStepSound     float32
}

func (msg *PlayerCommandPhysicalParameters) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, msg.Overweight)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.TransitionSpeed)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, msg.IsHeavyBreathing)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.WalkOverweight)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.SoundRadius)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, msg.BreathIsAudible)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, msg.MinStepSound)
	if err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandPhysicalParameters) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &msg.Overweight)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.TransitionSpeed)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &msg.IsHeavyBreathing)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.WalkOverweight)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.SoundRadius)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &msg.BreathIsAudible)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &msg.MinStepSound)
	if err != nil {
		return err
	}
	return nil
}
