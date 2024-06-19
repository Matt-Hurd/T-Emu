package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type LampControllerInfo struct {
	Id    int32
	State enums.ELampState
}

func (l *LampControllerInfo) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadInt32(buffer, &l.Id); err != nil {
		return err
	}
	var state byte
	if err := helpers.ReadByte(buffer, &state); err != nil {
		return err
	} else {
		l.State = enums.ELampState(state)
	}
	return nil
}

func (l *LampControllerInfo) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteInt32(buffer, l.Id); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(l.State)); err != nil {
		return err
	}
	return nil
}
