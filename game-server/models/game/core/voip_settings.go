package core

import (
	"bytes"
	"game-server/helpers"
)

type VoipSettings struct {
	VoipEnabled         bool
	VoipQualitySettings VoipQualitySettings
	PushToTalkSettings  PushToTalkSettings
}

func (vs *VoipSettings) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadBool(buffer, &vs.VoipEnabled); err != nil {
		return err
	}
	if vs.VoipQualitySettings, err = DeserializeVoipQualitySettings(buffer); err != nil {
		return err
	}
	if vs.PushToTalkSettings, err = DeserializePushToTalkSettings(buffer); err != nil {
		return err
	}
	return nil
}

func (vs *VoipSettings) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteBool(buffer, vs.VoipEnabled); err != nil {
		return err
	}
	if err = vs.VoipQualitySettings.Serialize(buffer); err != nil {
		return err
	}
	if err = vs.PushToTalkSettings.Serialize(buffer); err != nil {
		return err
	}
	return nil
}
