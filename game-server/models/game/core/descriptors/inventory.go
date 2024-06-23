package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type InventoryDescriptor struct {
	Equipment       ItemDescriptor
	Stash           *ItemDescriptor
	QuestRaidItems  *ItemDescriptor
	QuestStashItems *ItemDescriptor
	SortingTable    *ItemDescriptor
	FastAccess      *FastAccessDescriptor
	DiscardLimits   *DiscardLimitsDescriptor
}

func (inventory *InventoryDescriptor) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = inventory.Equipment.Serialize(buffer)
	if err != nil {
		return err
	}
	err = helpers.WriteBool(buffer, inventory.Stash != nil)
	if err != nil {
		return err
	}
	if inventory.Stash != nil {
		err = inventory.Stash.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, inventory.QuestRaidItems != nil)
	if err != nil {
		return err
	}
	if inventory.QuestRaidItems != nil {
		err = inventory.QuestRaidItems.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, inventory.QuestStashItems != nil)
	if err != nil {
		return err
	}
	if inventory.QuestStashItems != nil {
		err = inventory.QuestStashItems.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, inventory.SortingTable != nil)
	if err != nil {
		return err
	}
	if inventory.SortingTable != nil {
		err = inventory.SortingTable.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, inventory.FastAccess != nil)
	if err != nil {
		return err
	}
	if inventory.FastAccess != nil {
		err = inventory.FastAccess.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	err = helpers.WriteBool(buffer, inventory.DiscardLimits != nil)
	if err != nil {
		return err
	}
	if inventory.DiscardLimits != nil {
		err = inventory.DiscardLimits.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inventory *InventoryDescriptor) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = inventory.Equipment.Deserialize(buffer)
	if err != nil {
		return err
	}
	var hasStash bool
	err = helpers.ReadBool(buffer, &hasStash)
	if err != nil {
		return err
	}
	if hasStash {
		inventory.Stash = &ItemDescriptor{}
		err = inventory.Stash.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var hasQuestRaidItems bool
	err = helpers.ReadBool(buffer, &hasQuestRaidItems)
	if err != nil {
		return err
	}
	if hasQuestRaidItems {
		inventory.QuestRaidItems = &ItemDescriptor{}
		err = inventory.QuestRaidItems.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var hasQuestStashItems bool
	err = helpers.ReadBool(buffer, &hasQuestStashItems)
	if err != nil {
		return err
	}
	if hasQuestStashItems {
		inventory.QuestStashItems = &ItemDescriptor{}
		err = inventory.QuestStashItems.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var hasSortingTable bool
	err = helpers.ReadBool(buffer, &hasSortingTable)
	if err != nil {
		return err
	}
	if hasSortingTable {
		inventory.SortingTable = &ItemDescriptor{}
		err = inventory.SortingTable.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var hasFastAccess bool
	err = helpers.ReadBool(buffer, &hasFastAccess)
	if err != nil {
		return err
	}
	if hasFastAccess {
		inventory.FastAccess = &FastAccessDescriptor{}
		err = inventory.FastAccess.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	var hasDiscardLimits bool
	err = helpers.ReadBool(buffer, &hasDiscardLimits)
	if err != nil {
		return err
	}
	if hasDiscardLimits {
		inventory.DiscardLimits = &DiscardLimitsDescriptor{}
		err = inventory.DiscardLimits.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
