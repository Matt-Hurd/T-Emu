package descriptors

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type FireModeComponentDescriptor struct {
	FireMode enums.FireMode
}

func (fireModeComponent *FireModeComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, int32(fireModeComponent.FireMode)); err != nil {
		return err
	}
	return nil
}

func (fireModeComponent *FireModeComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var val int32
	if err := helpers.ReadInt32(buffer, &val); err != nil {
		return err
	}
	fireModeComponent.FireMode = enums.FireMode(val)
	return nil
}
