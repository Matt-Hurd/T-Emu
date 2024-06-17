package core

import (
	"bytes"
	"game-server/helpers"
)

type PushToTalkSettings struct {
	SpeakingSecondsLimit    float32
	SpeakingSecondsInterval float32
	ActivationsLimit        byte
	ActivationsInterval     float32
	BlockingTime            float32
	AlertDistanceMeters     byte
	HearingDistance         byte
	AbuseTraceSeconds       float32
}

func DeserializePushToTalkSettings(buffer *bytes.Buffer) (PushToTalkSettings, error) {
	var pts PushToTalkSettings
	var err error

	if err = helpers.ReadFloat32(buffer, &pts.SpeakingSecondsLimit); err != nil {
		return pts, err
	}
	if err = helpers.ReadFloat32(buffer, &pts.SpeakingSecondsInterval); err != nil {
		return pts, err
	}
	if err = helpers.ReadByte(buffer, &pts.ActivationsLimit); err != nil {
		return pts, err
	}
	if err = helpers.ReadFloat32(buffer, &pts.ActivationsInterval); err != nil {
		return pts, err
	}
	if err = helpers.ReadFloat32(buffer, &pts.BlockingTime); err != nil {
		return pts, err
	}
	if err = helpers.ReadByte(buffer, &pts.AlertDistanceMeters); err != nil {
		return pts, err
	}
	if err = helpers.ReadByte(buffer, &pts.HearingDistance); err != nil {
		return pts, err
	}
	if err = helpers.ReadFloat32(buffer, &pts.AbuseTraceSeconds); err != nil {
		return pts, err
	}
	return pts, nil
}

func (pts *PushToTalkSettings) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteFloat32(buffer, pts.SpeakingSecondsLimit); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, pts.SpeakingSecondsInterval); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, pts.ActivationsLimit); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, pts.ActivationsInterval); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, pts.BlockingTime); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, pts.AlertDistanceMeters); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, pts.HearingDistance); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, pts.AbuseTraceSeconds); err != nil {
		return err
	}
	return nil
}
