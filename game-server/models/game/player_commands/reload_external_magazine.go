package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type PlayerCommandReloadExternalMagazine struct {
	AmmoInChamberCurrent   int32
	AmmoInChamberResult    int32
	NeedToAddAmmoInChamber bool
	ShellsInWeapon         int32
	WeaponLevel            int32
	InInventory            bool
	MagInWeapon            bool
	AmmoInMag              int32
	Malfunction            bool
	MagOutFromInventory    bool
	SlotModeID             string
	NextMag                core.ComponentialItem
	IsQuickReload          bool
	OldMagAnimIndex        int32
	NewMagAnimIndex        int32
	InMisfireMalfunction   bool
}

func (p *PlayerCommandReloadExternalMagazine) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, p.AmmoInChamberCurrent); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.AmmoInChamberResult); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.NeedToAddAmmoInChamber); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.ShellsInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.WeaponLevel); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.InInventory); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.MagInWeapon); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.AmmoInMag); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.Malfunction); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.MagOutFromInventory); err != nil {
		return err
	}
	if err := helpers.WriteString(buffer, p.SlotModeID); err != nil {
		return err
	}
	if err := p.NextMag.Serialize(buffer); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.IsQuickReload); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.OldMagAnimIndex); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, p.NewMagAnimIndex); err != nil {
		return err
	}
	if err := helpers.WriteBool(buffer, p.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}

func (p *PlayerCommandReloadExternalMagazine) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadInt32(buffer, &p.AmmoInChamberCurrent); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.AmmoInChamberResult); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.NeedToAddAmmoInChamber); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.ShellsInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.WeaponLevel); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.InInventory); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.MagInWeapon); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.AmmoInMag); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.Malfunction); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.MagOutFromInventory); err != nil {
		return err
	}
	if err = helpers.ReadString(buffer, &p.SlotModeID); err != nil {
		return err
	}
	if err = p.NextMag.Deserialize(buffer); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.IsQuickReload); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.OldMagAnimIndex); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, &p.NewMagAnimIndex); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &p.InMisfireMalfunction); err != nil {
		return err
	}
	return nil
}
