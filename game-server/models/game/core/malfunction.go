package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1505
type Malfunction struct {
	Malfunction                            byte
	LastShotOverheat                       float32
	LastShotTime                           float32
	SlideOnOverheatReached                 bool
	PlayersWhoKnowAboutMalfunction         []string
	PlayersWhoKnowMalfType                 []string
	PlayersReducedMalfChances              map[string]byte
	hasAmmoToFireTemplateId                bool
	AmmoToFireTemplateId                   string
	hasAmmoWillBeLoadedToChamberTemplateId bool
	AmmoWillBeLoadedToChamberTemplateId    string
	hasAmmoMalfunctionedTemplateId         bool
	AmmoMalfunctionedTemplateId            string
}

func (malfunction *Malfunction) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteByte(buffer, malfunction.Malfunction)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, malfunction.LastShotOverheat)
	if err != nil {
		return err
	}
	err = helpers.WriteFloat32(buffer, malfunction.LastShotTime)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, malfunction.SlideOnOverheatReached)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(len(malfunction.PlayersWhoKnowAboutMalfunction)))
	if err != nil {
		return err
	}
	for _, v := range malfunction.PlayersWhoKnowAboutMalfunction {
		err = helpers.WriteUTF16String(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(malfunction.PlayersWhoKnowMalfType)))
	if err != nil {
		return err
	}
	for _, v := range malfunction.PlayersWhoKnowMalfType {
		err = helpers.WriteUTF16String(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteInt32(buffer, int32(len(malfunction.PlayersReducedMalfChances)))
	if err != nil {
		return err
	}
	for k, v := range malfunction.PlayersReducedMalfChances {
		err = helpers.WriteUTF16String(buffer, k)
		if err != nil {
			return err
		}
		err = helpers.WriteByte(buffer, v)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, malfunction.hasAmmoToFireTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoToFireTemplateId {
		err = helpers.WriteUTF16String(buffer, malfunction.AmmoToFireTemplateId)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, malfunction.hasAmmoWillBeLoadedToChamberTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoWillBeLoadedToChamberTemplateId {
		err = helpers.WriteUTF16String(buffer, malfunction.AmmoWillBeLoadedToChamberTemplateId)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, malfunction.hasAmmoMalfunctionedTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoMalfunctionedTemplateId {
		err = helpers.WriteUTF16String(buffer, malfunction.AmmoMalfunctionedTemplateId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (malfunction *Malfunction) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadByte(buffer, &malfunction.Malfunction)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &malfunction.LastShotOverheat)
	if err != nil {
		return err
	}
	err = helpers.ReadFloat32(buffer, &malfunction.LastShotTime)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &malfunction.SlideOnOverheatReached)
	if err != nil {
		return err
	}
	var PlayersWhoKnowAboutMalfunctionLength int32
	err = helpers.ReadInt32(buffer, &PlayersWhoKnowAboutMalfunctionLength)
	if err != nil {
		return err
	}
	malfunction.PlayersWhoKnowAboutMalfunction = make([]string, PlayersWhoKnowAboutMalfunctionLength)
	for i := range malfunction.PlayersWhoKnowAboutMalfunction {
		err = helpers.ReadUTF16String(buffer, &malfunction.PlayersWhoKnowAboutMalfunction[i])
		if err != nil {
			return err
		}
	}
	var PlayersWhoKnowMalfTypeLength int32
	err = helpers.ReadInt32(buffer, &PlayersWhoKnowMalfTypeLength)
	if err != nil {
		return err
	}
	malfunction.PlayersWhoKnowMalfType = make([]string, PlayersWhoKnowMalfTypeLength)
	for i := range malfunction.PlayersWhoKnowMalfType {
		err = helpers.ReadUTF16String(buffer, &malfunction.PlayersWhoKnowMalfType[i])
		if err != nil {
			return err
		}
	}
	var PlayersReducedMalfChancesLength int32
	err = helpers.ReadInt32(buffer, &PlayersReducedMalfChancesLength)
	if err != nil {
		return err
	}
	malfunction.PlayersReducedMalfChances = make(map[string]byte)
	for i := 0; i < int(PlayersReducedMalfChancesLength); i++ {
		var key string
		err = helpers.ReadUTF16String(buffer, &key)
		if err != nil {
			return err
		}
		var value byte
		err = helpers.ReadByte(buffer, &value)
		if err != nil {
			return err
		}
		malfunction.PlayersReducedMalfChances[key] = value
	}
	err = helpers.ReadBool(buffer, &malfunction.hasAmmoToFireTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoToFireTemplateId {
		err = helpers.ReadUTF16String(buffer, &malfunction.AmmoToFireTemplateId)
		if err != nil {
			return err
		}
	}
	err = helpers.ReadBool(buffer, &malfunction.hasAmmoWillBeLoadedToChamberTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoWillBeLoadedToChamberTemplateId {
		err = helpers.ReadUTF16String(buffer, &malfunction.AmmoWillBeLoadedToChamberTemplateId)
		if err != nil {
			return err
		}
	}
	err = helpers.ReadBool(buffer, &malfunction.hasAmmoMalfunctionedTemplateId)
	if err != nil {
		return err
	}
	if malfunction.hasAmmoMalfunctionedTemplateId {
		err = helpers.ReadUTF16String(buffer, &malfunction.AmmoMalfunctionedTemplateId)
		if err != nil {
			return err
		}
	}
	return nil
}
