package core

import (
	"bytes"
	"game-server/helpers"
)

type WindowBreakerInfo struct {
	Id       int32
	Position Vector3
}

func (w *WindowBreakerInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadInt32(buffer, &w.Id); err != nil {
		return err
	}
	if err = w.Position.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}

func (w *WindowBreakerInfo) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteInt32(buffer, w.Id); err != nil {
		return err
	}
	if err = w.Position.Serialize(buffer); err != nil {
		return err
	}
	return nil
}
