package core

import (
	"bytes"
	"game-server/helpers"
	"time"
)

type GameDateTime struct {
	GameOnly     bool
	RealDateTime time.Time
	GameDateTime time.Time
	TimeFactor   float32
}

func DeserializeGameDateTime(buffer *bytes.Buffer) (GameDateTime, error) {
	var gdt GameDateTime
	var err error

	if err = helpers.ReadBool(buffer, &gdt.GameOnly); err != nil {
		return gdt, err
	}
	if !gdt.GameOnly {
		if err = helpers.ReadDateTime(buffer, &gdt.RealDateTime); err != nil {
			return gdt, err
		}
	}
	// buf := buffer.Next(8)
	// fmt.Printf("GameDateTime: %x\n", buf)
	if err = helpers.ReadDateTime(buffer, &gdt.GameDateTime); err != nil {
		return gdt, err
	}
	if err = helpers.ReadFloat32(buffer, &gdt.TimeFactor); err != nil {
		return gdt, err
	}
	return gdt, nil
}

func (gdt *GameDateTime) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteBool(buffer, gdt.GameOnly); err != nil {
		return err
	}
	if !gdt.GameOnly {
		if err = helpers.WriteDateTime(buffer, gdt.RealDateTime); err != nil {
			return err
		}
	}
	if err = helpers.WriteDateTime(buffer, gdt.GameDateTime); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, gdt.TimeFactor); err != nil {
		return err
	}
	return nil
}
