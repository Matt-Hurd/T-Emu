package playercommands

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type PlayerCommandMalfunction struct {
	MalfunctionState          enums.MalfunctionState
	AmmoToFire                string
	MalfunctionedAmmo         string
	AmmoWillBeLoadedToChamber string
}

func (msg *PlayerCommandMalfunction) Serialize(buffer *bytes.Buffer) error {
	if err := buffer.WriteByte(byte(msg.MalfunctionState)); err != nil {
		return err
	}
	if msg.MalfunctionState == enums.MalfunctionStateNone {
		return nil
	}
	if err := helpers.WriteBool(buffer, len(msg.AmmoToFire) > 0); err != nil {
		return err
	}
	if len(msg.AmmoToFire) > 0 {
		if err := helpers.WriteMongoId(buffer, msg.AmmoToFire); err != nil {
			return err
		}
	}
	if err := helpers.WriteBool(buffer, len(msg.MalfunctionedAmmo) > 0); err != nil {
		return err
	}
	if len(msg.MalfunctionedAmmo) > 0 {
		if err := helpers.WriteMongoId(buffer, msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	if err := helpers.WriteBool(buffer, len(msg.AmmoWillBeLoadedToChamber) > 0); err != nil {
		return err
	}
	if len(msg.AmmoWillBeLoadedToChamber) > 0 {
		if err := helpers.WriteMongoId(buffer, msg.AmmoWillBeLoadedToChamber); err != nil {
			return err
		}
	}
	return nil
}

func (msg *PlayerCommandMalfunction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, (*byte)(&msg.MalfunctionState)); err != nil {
		return err
	}
	if msg.MalfunctionState == enums.MalfunctionStateNone {
		return nil
	}
	var hasAmmoToFire bool
	if err = helpers.ReadBool(buffer, &hasAmmoToFire); err != nil {
		return err
	}
	if hasAmmoToFire {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoToFire); err != nil {
			return err
		}
	}
	var hasMalfunctionedAmmo bool
	if err = helpers.ReadBool(buffer, &hasMalfunctionedAmmo); err != nil {
		return err
	}
	if hasMalfunctionedAmmo {
		if err = helpers.ReadMongoId(buffer, &msg.MalfunctionedAmmo); err != nil {
			return err
		}
	}
	var hasAmmoWillBeLoadedToChamber bool
	if err = helpers.ReadBool(buffer, &hasAmmoWillBeLoadedToChamber); err != nil {
		return err
	}
	if hasAmmoWillBeLoadedToChamber {
		if err = helpers.ReadMongoId(buffer, &msg.AmmoWillBeLoadedToChamber); err != nil {
			return err
		}
	}
	return nil
}
