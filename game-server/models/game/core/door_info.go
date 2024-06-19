package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type DoorInfo struct {
	NetId    int16
	State    enums.EDoorState
	IsBroken bool
}

func (d *DoorInfo) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadInt16(buffer, &d.NetId); err != nil {
		return err
	}
	var stateInfo byte
	if err = helpers.ReadByte(buffer, &stateInfo); err != nil {
		return err
	}
	d.State = enums.EDoorState(stateInfo & 239)
	d.IsBroken = (stateInfo & 16) > 0
	return nil
}

func (d *DoorInfo) Serialize(buffer *bytes.Buffer) error {
	var err error
	var stateInfo byte

	if err = helpers.WriteInt16(buffer, d.NetId); err != nil {
		return err
	}
	stateInfo = byte(d.State) | (helpers.BoolToByte(d.IsBroken) << 4)
	if err = helpers.WriteByte(buffer, stateInfo); err != nil {
		return err
	}
	return nil
}
