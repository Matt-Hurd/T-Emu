package core

import (
	"bytes"
	"game-server/helpers"
)

type ExfiltrationController struct {
	ExfilDatas []ExfilData
}

func (e *ExfiltrationController) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var num int16

	if err = helpers.ReadInt16(buffer, &num); err != nil {
		return err
	}
	e.ExfilDatas = make([]ExfilData, num)
	for i := 0; i < int(num); i++ {
		var exfilData ExfilData
		if err = exfilData.Deserialize(buffer); err != nil {
			return err
		}
		e.ExfilDatas[i] = exfilData
	}
	return nil
}

func (e *ExfiltrationController) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteInt16(buffer, int16(len(e.ExfilDatas))); err != nil {
		return err
	}
	for _, exfilData := range e.ExfilDatas {
		if err = exfilData.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}
