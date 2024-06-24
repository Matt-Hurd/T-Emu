package descriptors

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

type DogTagComponentDescriptor struct {
	AccountId            string
	ProfileId            string
	Nickname             string
	Side                 enums.PlayerSide
	Level                int32
	Time                 float64
	Status               string
	KillerAccountId      string
	KillerProfileId      string
	KillerName           string
	WeaponName           string
	CarriedByGroupMember bool
}

func (dogTagComponent *DogTagComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, dogTagComponent.AccountId)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.ProfileId)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.Nickname)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(dogTagComponent.Side))
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, dogTagComponent.Level)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat64(buffer, dogTagComponent.Time)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.Status)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.KillerAccountId)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.KillerProfileId)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.KillerName)
	if err != nil {
		return err
	}
	err = helpers.WriteUTF16String(buffer, dogTagComponent.WeaponName)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, dogTagComponent.CarriedByGroupMember)
	if err != nil {
		return err
	}
	return nil
}

func (dogTagComponent *DogTagComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.AccountId)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.ProfileId)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.Nickname)
	if err != nil {
		return err
	}
	var side int32
	err = helpers.ReadInt32(buffer, &side)
	if err != nil {
		return err
	}
	dogTagComponent.Side = enums.PlayerSide(side)
	err = helpers.ReadInt32(buffer, &dogTagComponent.Level)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat64(buffer, &dogTagComponent.Time)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.Status)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.KillerAccountId)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.KillerProfileId)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.KillerName)
	if err != nil {
		return err
	}
	err = helpers.ReadUTF16String(buffer, &dogTagComponent.WeaponName)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &dogTagComponent.CarriedByGroupMember)
	if err != nil {
		return err
	}
	return nil
}
