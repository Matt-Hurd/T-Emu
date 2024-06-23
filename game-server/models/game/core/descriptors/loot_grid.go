package descriptors

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/enums"
)

// GClass1507
type LootGridDescriptor struct {
	Id             string
	ContainedItems []GridItemDescriptor
}

// GClass1506
type GridItemDescriptor struct {
	Location LocationInGrid
	Item     ItemDescriptor
}

type LocationInGrid struct {
	X          int32
	Y          int32
	Rotation   enums.ItemRotation
	IsSearched bool
}

func (lig *LocationInGrid) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteInt32(buffer, lig.X)
	if err != nil {
		return err
	}
	err = helpers.WriteInt32(buffer, lig.Y)
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, byte(lig.Rotation))
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, lig.IsSearched)
	if err != nil {
		return err
	}
	return nil
}

func (lig *LocationInGrid) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadInt32(buffer, &lig.X)
	if err != nil {
		return err
	}
	err = helpers.ReadInt32(buffer, &lig.Y)
	if err != nil {
		return err
	}
	var rotation byte
	err = helpers.ReadByte(buffer, &rotation)
	if err != nil {
		return err
	}
	lig.Rotation = enums.ItemRotation(rotation)
	err = helpers.ReadBool(buffer, &lig.IsSearched)
	if err != nil {
		return err
	}
	return nil
}

func (lg *LootGridDescriptor) Serialize(buffer *bytes.Buffer) error {
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
		err = v.Location.Serialize(buffer)
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

func (lg *LootGridDescriptor) Deserialize(buffer *bytes.Buffer) error {
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
		var item GridItemDescriptor
		err = item.Location.Deserialize(buffer)
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
