package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandDoorUnlockInteraction struct {
	InteractionDoor       string
	InteractionDoorKey    string
	InteractionDoorResult bool
}

func (msg *PlayerCommandDoorUnlockInteraction) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.InteractionDoor); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.InteractionDoorKey); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InteractionDoorResult); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDoorUnlockInteraction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.InteractionDoor); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.InteractionDoorKey); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InteractionDoorResult); err != nil {
		return err
	}
	return nil
}
