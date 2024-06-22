package core

import (
	"game-server/helpers"
	"game-server/models/game/enums"
)

// GClass1405
type SearchLoot struct {
	Flag   bool
	Id     string
	States map[string]enums.SearchedState
	Grids  []GClass1406
}

// GClass1406
type GClass1406 struct {
	Flag       bool
	Id         string
	KnownItems map[string][]string
}

func (gClass1406 *GClass1406) Serialize(packedWriter *helpers.LimitedWriter) error {
	err := packedWriter.Write(gClass1406.Flag)
	if err != nil {
		return err
	}
	if gClass1406.Flag {
		return nil
	}
	err = packedWriter.WriteLimitedString(gClass1406.Id, ' ', 'z')
	if err != nil {
		return err
	}
	err = packedWriter.Write(int32(len(gClass1406.KnownItems)))
	if err != nil {
		return err
	}
	for k, v := range gClass1406.KnownItems {
		err = packedWriter.WriteLimitedString(k, ' ', 'z')
		if err != nil {
			return err
		}
		err = packedWriter.Write(int32(len(v)))
		if err != nil {
			return err
		}
		for _, v2 := range v {
			err = packedWriter.WriteLimitedString(v2, ' ', 'z')
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (gClass1406 *GClass1406) Deserialize(packedReader *helpers.LimitedReader) error {
	var err error
	err = packedReader.Read(&gClass1406.Flag)
	if err != nil {
		return err
	}
	if gClass1406.Flag {
		return nil
	}
	gClass1406.Id, err = packedReader.ReadLimitedString(' ', 'z')
	if err != nil {
		return err
	}
	var KnownItemsLength int32
	err = packedReader.Read(&KnownItemsLength)
	if err != nil {
		return err
	}
	gClass1406.KnownItems = make(map[string][]string, KnownItemsLength)
	for i := int32(0); i < KnownItemsLength; i++ {
		key, err := packedReader.ReadLimitedString(' ', 'z')
		if err != nil {
			return err
		}
		var valueLength int32
		err = packedReader.Read(&valueLength)
		if err != nil {
			return err
		}
		gClass1406.KnownItems[key] = make([]string, valueLength)
		for j := int32(0); j < valueLength; j++ {
			gClass1406.KnownItems[key][j], err = packedReader.ReadLimitedString(' ', 'z')
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (searchLoot *SearchLoot) Serialize(packedWriter *helpers.LimitedWriter) error {
	err := packedWriter.Write(searchLoot.Flag)
	if err != nil {
		return err
	}
	if searchLoot.Flag {
		return nil
	}

	err = packedWriter.WriteLimitedString(searchLoot.Id, ' ', 'z')
	if err != nil {
		return err
	}
	err = packedWriter.Write(int32(len(searchLoot.States)))
	if err != nil {
		return err
	}
	for k, v := range searchLoot.States {
		err = packedWriter.WriteLimitedString(k, ' ', 'z')
		if err != nil {
			return err
		}
		err = packedWriter.WriteEnum(int(enums.Unsearched), int(enums.FullySearched), int(v))
		if err != nil {
			return err
		}
	}
	err = packedWriter.Write(int32(len(searchLoot.Grids)))
	if err != nil {
		return err
	}
	for _, v := range searchLoot.Grids {
		err = v.Serialize(packedWriter)
		if err != nil {
			return err
		}
	}
	return nil
}

func (searchLoot *SearchLoot) Deserialize(packedReader *helpers.LimitedReader) error {
	var err error
	err = packedReader.Read(&searchLoot.Flag)
	if err != nil {
		return err
	}
	if searchLoot.Flag {
		return nil
	}

	searchLoot.Id, err = packedReader.ReadLimitedString(' ', 'z')
	if err != nil {
		return err
	}
	var StatesLength int32
	err = packedReader.Read(&StatesLength)
	if err != nil {
		return err
	}
	searchLoot.States = make(map[string]enums.SearchedState, StatesLength)
	for i := int32(0); i < StatesLength; i++ {
		key, err := packedReader.ReadLimitedString(' ', 'z')
		if err != nil {
			return err
		}
		val, err := packedReader.ReadEnum(int(enums.Unsearched), int(enums.FullySearched))
		if err != nil {
			return err
		}
		searchLoot.States[key] = enums.SearchedState(val)
	}
	var GridsLength int32
	err = packedReader.Read(&GridsLength)
	if err != nil {
		return err
	}
	searchLoot.Grids = make([]GClass1406, GridsLength)
	for i := int32(0); i < GridsLength; i++ {
		err = searchLoot.Grids[i].Deserialize(packedReader)
		if err != nil {
			return err
		}
	}
	return nil
}
