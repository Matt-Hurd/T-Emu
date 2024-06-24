package descriptors

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/math"
)

// GClass1532
type LootDescriptor struct {
	Serializable
	hasID          bool
	Id             string
	Position       math.Vector3
	Rotation       math.Vector3
	Item           ItemDescriptor
	hasProfiles    bool
	ValidProfiles  []string
	IsContainer    bool
	UseGravity     bool
	RandomRotation bool
	Shift          math.Vector3
	PlatformId     int16
}

func (loot *LootDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteBool(buffer, loot.hasID)
	if err != nil {
		return err
	}
	if loot.hasID {
		err = helpers.WriteUTF16String(buffer, loot.Id)
		if err != nil {
			return err
		}
	}
	err = loot.Position.Serialize(buffer)
	if err != nil {
		return err
	}
	err = loot.Rotation.Serialize(buffer)
	if err != nil {
		return err
	}
	err = loot.Item.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, loot.hasProfiles)
	if err != nil {
		return err
	}
	if loot.hasProfiles {
		for _, v := range loot.ValidProfiles {
			err = helpers.WriteUTF16String(buffer, v)
			if err != nil {
				return err
			}
		}
	}
	err = helpers.WriteBool(buffer, loot.IsContainer)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, loot.UseGravity)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, loot.RandomRotation)
	if err != nil {
		return err
	}
	err = loot.Shift.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteInt16(buffer, loot.PlatformId)
	if err != nil {
		return err
	}
	return nil
}

func (loot *LootDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadBool(buffer, &loot.hasID)
	if err != nil {
		return err
	}
	if loot.hasID {
		err = helpers.ReadUTF16String(buffer, &loot.Id)
		if err != nil {
			return err
		}
	}
	err = loot.Position.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = loot.Rotation.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = loot.Item.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &loot.hasProfiles)
	if err != nil {
		return err
	}
	if loot.hasProfiles {
		loot.ValidProfiles = make([]string, 0)
		var validProfilesLength int32
		err = helpers.ReadInt32(buffer, &validProfilesLength)
		if err != nil {
			return err
		}
		for i := int32(0); i < validProfilesLength; i++ {
			var validProfile string
			err = helpers.ReadString(buffer, &validProfile)
			if err != nil {
				return err
			}
			loot.ValidProfiles = append(loot.ValidProfiles, validProfile)
		}
	}
	err = helpers.ReadBool(buffer, &loot.IsContainer)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &loot.UseGravity)
	if err != nil {
		return err
	}
	err = helpers.ReadBool(buffer, &loot.RandomRotation)
	if err != nil {
		return err
	}
	err = loot.Shift.Deserialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.ReadInt16(buffer, &loot.PlatformId)
	if err != nil {
		return err
	}
	return nil
}
