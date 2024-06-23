package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PlayerCommandModChanged struct {
	ModAdded   bool
	ItemID     string
	SlotModeID string
	ModeItems  core.ComponentialItem
}

func (msg *PlayerCommandModChanged) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteBool(buffer, msg.ModAdded); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.ItemID); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.SlotModeID); err != nil {
		return err
	}
	if msg.ModAdded {
		if err := msg.ModeItems.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandModChanged) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadBool(buffer, &msg.ModAdded); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.ItemID); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.SlotModeID); err != nil {
		return err
	}
	if msg.ModAdded {
		if err = msg.ModeItems.Deserialize(buffer); err != nil {
			return err
		}
	}
	return nil
}
