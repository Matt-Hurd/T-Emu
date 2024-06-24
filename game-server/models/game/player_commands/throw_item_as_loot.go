package playercommands

import (
	"bytes"
	"game-server/models/game/core"
	"game-server/models/game/math"
)

type PlayerCommandThrowItemAsLoot struct {
	Position        math.Vector3
	Rotation        math.Quaternion
	Velocity        math.Vector3
	AngularVelocity math.Vector3
	Item            core.ComponentialItem
}

func (msg *PlayerCommandThrowItemAsLoot) Serialize(buffer *bytes.Buffer) error {
	if err := msg.Position.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.Rotation.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.Velocity.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.AngularVelocity.Serialize(buffer); err != nil {
		return err
	}
	if err := msg.Item.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandThrowItemAsLoot) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = msg.Position.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.Rotation.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.Velocity.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.AngularVelocity.Deserialize(buffer); err != nil {
		return err
	}
	if err = msg.Item.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
