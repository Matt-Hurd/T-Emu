package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandDoorBreachInteraction struct {
	InteractionDoor       string
	InteractionDoorResult bool
}

func (msg *PlayerCommandDoorBreachInteraction) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.InteractionDoor); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, msg.InteractionDoorResult); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandDoorBreachInteraction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.InteractionDoor); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &msg.InteractionDoorResult); err != nil {
		return err
	}
	return nil
}
