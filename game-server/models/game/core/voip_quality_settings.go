package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type VoipQualitySettings struct {
	FrameSize              enums.FrameSize
	AudioQuality           enums.AudioQuality
	ForwardErrorCorrection bool
	NoiseSuppression       enums.NoiseSuppressionLevels
	SensitivityLevels      enums.VadSensitivityLevels
}

func DeserializeVoipQualitySettings(buffer *bytes.Buffer) (VoipQualitySettings, error) {
	var vqs VoipQualitySettings
	var err error
	var val byte

	if err = helpers.ReadByte(buffer, &val); err != nil {
		return vqs, err
	} else {
		vqs.FrameSize = enums.FrameSize(val)
	}
	if err = helpers.ReadByte(buffer, &val); err != nil {
		return vqs, err
	} else {
		vqs.AudioQuality = enums.AudioQuality(val)
	}
	if err = helpers.ReadBool(buffer, &vqs.ForwardErrorCorrection); err != nil {
		return vqs, err
	}
	if err = helpers.ReadByte(buffer, &val); err != nil {
		return vqs, err
	} else {
		vqs.NoiseSuppression = enums.NoiseSuppressionLevels(val)
	}
	if err = helpers.ReadByte(buffer, &val); err != nil {
		return vqs, err
	} else {
		vqs.SensitivityLevels = enums.VadSensitivityLevels(val)
	}
	return vqs, nil
}

func (vqs *VoipQualitySettings) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteByte(buffer, byte(vqs.FrameSize)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(vqs.AudioQuality)); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, vqs.ForwardErrorCorrection); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(vqs.NoiseSuppression)); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(vqs.SensitivityLevels)); err != nil {
		return err
	}
	return nil
}
