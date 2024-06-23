package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PlayerCommandSyncStationaryMagazine struct {
	StationaryWeaponID string
	SlotModeID         string
	Items              core.ComponentialItem
}

func (msg *PlayerCommandSyncStationaryMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.StationaryWeaponID); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, msg.SlotModeID); err != nil {
		return err
	}
	if err := msg.Items.Serialize(buffer); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandSyncStationaryMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.StationaryWeaponID); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &msg.SlotModeID); err != nil {
		return err
	}
	if err = msg.Items.Deserialize(buffer); err != nil {
		return err
	}
	return nil
}
