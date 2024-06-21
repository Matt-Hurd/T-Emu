package core

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

// GClass1507
type LootGrid struct {
	Id             string
	ContainedItems []GridItem
}

// GClass1506
type GridItem struct {
	Location LocationInGrid
	Item     Item
}

type LocationInGrid struct {
	X          int32
	Y          int32
	Rotation   enums.ItemRotation
	IsSearched bool
}

func (lg *LootGrid) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteUTF16String(buffer, lg.Id)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, int32(len(lg.ContainedItems)))
	if err != nil {
		return err
	}
	for _, v := range lg.ContainedItems {
		err = helpers.WriteInt32(buffer, v.Location.X)
		if err != nil {
			return err
		}
		err = helpers.WriteInt32(buffer, v.Location.Y)
		if err != nil {
			return err
		}
		err = helpers.WriteByte(buffer, byte(v.Location.Rotation))
		if err != nil {
			return err
		}
		err = helpers.WriteBool(buffer, v.Location.IsSearched)
		if err != nil {
			return err
		}
		err = v.Item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lg *LootGrid) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadUTF16String(buffer, &lg.Id)
	if err != nil {
		return err
	}
	var count int32
	err = helpers.ReadInt32(buffer, &count)
	if err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {
		var item GridItem
		err = helpers.ReadInt32(buffer, &item.Location.X)
		if err != nil {
			return err
		}
		err = helpers.ReadInt32(buffer, &item.Location.Y)
		if err != nil {
			return err
		}
		var rotation byte
		err = helpers.ReadByte(buffer, &rotation)
		if err != nil {
			return err
		}
		item.Location.Rotation = enums.ItemRotation(rotation)
		err = helpers.ReadBool(buffer, &item.Location.IsSearched)
		if err != nil {
			return err
		}
		err = item.Item.Deserialize(buffer)
		if err != nil {
			return err
		}
		lg.ContainedItems = append(lg.ContainedItems, item)
	}
	return nil
}
