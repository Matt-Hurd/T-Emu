package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type HealthEffectType byte

const (
	HealthEffectTypeNone HealthEffectType = iota
	HealthEffectTypeHeavyBleeding
	HealthEffectTypeFracture
)

type PlayerCommandEffectOnPlayer struct {
	ChangedEffectType HealthEffectType
	IsActive          bool
}

func (msg *PlayerCommandEffectOnPlayer) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.ChangedEffectType)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.IsActive); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandEffectOnPlayer) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.ChangedEffectType)); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.IsActive); err != nil {
		return err
	}
	return nil
}
