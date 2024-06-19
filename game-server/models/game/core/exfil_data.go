package core

import (
	"bytes"
	"game-server/helpers"
)

type EExfiltrationStatus byte

const (
	NotPresent             EExfiltrationStatus = 1
	UncompleteRequirements EExfiltrationStatus = 2
	Countdown              EExfiltrationStatus = 3
	RegularMode            EExfiltrationStatus = 4
	Pending                EExfiltrationStatus = 5
	AwaitsManualActivation EExfiltrationStatus = 6
)

type ExfilData struct {
	Name               string
	ExfiltrationStatus EExfiltrationStatus
	StartTime          int32
	PlayerIds          []string
}

func (e *ExfilData) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var status byte

	if err = helpers.ReadString(buffer, &e.Name); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, &status); err != nil {
		return err
	} else {
		e.ExfiltrationStatus = EExfiltrationStatus(status)
	}
	if err = helpers.ReadInt32(buffer, &e.StartTime); err != nil {
		return err
	}
	if e.ExfiltrationStatus == Countdown {
		var startTime int16
		if err = helpers.ReadInt16(buffer, &startTime); err != nil {
			return err
		} else {
			e.StartTime = int32(startTime)
		}
	}
	var playerIdCounts int16
	if err = helpers.ReadInt16(buffer, &playerIdCounts); err != nil {
		return err
	}
	e.PlayerIds = make([]string, playerIdCounts)
	for i := 0; i < int(playerIdCounts); i++ {
		if err = helpers.ReadString(buffer, &e.PlayerIds[i]); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExfilData) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteString(buffer, e.Name); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(e.ExfiltrationStatus)); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, e.StartTime); err != nil {
		return err
	}
	if e.ExfiltrationStatus == Countdown {
		if err = helpers.WriteInt16(buffer, int16(e.StartTime)); err != nil {
			return err
		}
	}
	if err = helpers.WriteInt16(buffer, int16(len(e.PlayerIds))); err != nil {
		return err
	}
	for _, playerId := range e.PlayerIds {
		if err = helpers.WriteString(buffer, playerId); err != nil {
			return err
		}
	}
	return nil
}
