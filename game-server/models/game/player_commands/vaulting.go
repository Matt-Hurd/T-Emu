package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
	"game-server/models/game/math"
)

type PlayerCommandVaulting struct {
	VaultingPoint           math.Vector3
	VaultingStrategy        enums.VaultingStrategy
	VaultingSpeed           float32
	VaultingHeight          float32
	VaultingLength          float32
	BehindObstacleRatio     float32
	AbsoluteForwardVelocity float32
}

func (msg *PlayerCommandVaulting) Serialize(buffer *bytes.Buffer) error {
	if err := msg.VaultingPoint.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteByte(buffer, byte(msg.VaultingStrategy)); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.VaultingSpeed); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.VaultingHeight); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.VaultingLength); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.BehindObstacleRatio); err != nil {
		return err
	}
	if err := helpers.WriteFloat32(buffer, msg.AbsoluteForwardVelocity); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandVaulting) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = msg.VaultingPoint.Deserialize(buffer); err != nil {
		return err
	}
	var VaultingStrategy byte
	if err = helpers.ReadByte(buffer, &VaultingStrategy); err != nil {
		return err
	}
	msg.VaultingStrategy = enums.VaultingStrategy(VaultingStrategy)
	if err = helpers.ReadFloat32(buffer, &msg.VaultingSpeed); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.VaultingHeight); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.VaultingLength); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.BehindObstacleRatio); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &msg.AbsoluteForwardVelocity); err != nil {
		return err
	}
	return nil
}
