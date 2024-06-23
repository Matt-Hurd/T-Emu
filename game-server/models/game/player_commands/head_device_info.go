package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type HeadDeviceType byte

const (
	FaceShield HeadDeviceType = iota
	NightVision
	ThermalVision
)

type PlayerCommandHeadDeviceInfo struct {
	DeviceType HeadDeviceType
	IsActive   bool
}

func (msg *PlayerCommandHeadDeviceInfo) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.DeviceType)); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.IsActive); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandHeadDeviceInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.DeviceType)); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.IsActive); err != nil {
		return err
	}
	return nil
}
