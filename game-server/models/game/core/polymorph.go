package core

import (
	"bytes"
	"fmt"
	"game-server/helpers"
)

type Serializable interface {
	Deserialize(buffer *bytes.Buffer) error
	Serialize(buffer *bytes.Buffer) error
}

func ReadPolymorph(buffer *bytes.Buffer, o *Serializable) error {
	var t byte
	err := helpers.ReadByte(buffer, &t)
	if err != nil {
		return err
	}
	switch t {
	case 0:
		*o = &Quaternion{}
	case 1:
	// 	*o = &ClassTransformSync{}
	// case 2:
	// 	*o = &Vector3{}
	// case 3:
	// 	*o = &LocationInGrid{}
	// case 4:
	// 	*o = &WeightedLootPointSpawnPosition{}
	// case 5:
	// 	*o = &Inventory{}
	// case 6:
	// 	*o = &FastAccess{}
	// case 7:
	// 	*o = &DiscardLimits{}
	// case 8:
	// 	*o = &Slot{}
	// case 9:
	// 	*o = &ShellTemplate{}
	// case 10:
	// 	*o = &Malfunction{}
	// case 11:
	// 	*o = &ItemInGrid{}
	// case 12:
	// 	*o = &Grid{}
	// case 13:
	// 	*o = &StackSlot{}
	// case 14:
	// 	*o = &NestedItem{}
	// case 15:
	// 	*o = &Item{}
	case 17:
		*o = &FoodDrinkComponent{}
	// case 18:
	// 	*o = &PoisonComponent{}
	case 19:
		*o = &ResourceItemComponent{}
	// case 20:
	// 	*o = &LightComponent{}
	// case 21:
	// 	*o = &LockableComponent{}
	// case 22:
	// 	*o = &MapComponent{}
	case 23:
		*o = &MedKitComponent{}
	// case 24:
	// 	*o = &RepairableComponent{}
	// case 25:
	// 	*o = &SightComponent{}
	// case 26:
	// 	*o = &TogglableComponent{}
	// case 27:
	// 	*o = &FaceShieldComponent{}
	// case 28:
	// 	*o = &FoldableComponent{}
	// case 29:
	// 	*o = &FireModeComponent{}
	// case 30:
	// 	*o = &DogTagComponent{}
	// case 31:
	// 	*o = &TagComponent{}
	// case 32:
	// 	*o = &KeyComponent{}
	// case 33:
	// 	*o = &RepairKitComponent{}
	// case 34:
	// 	*o = &RepairEnhancementComponent{}
	// case 35:
	// 	*o = &RecodableComponent{}
	// case 36:
	// 	*o = &CultistAmuletComponent{}
	case 37:
		*o = &Loot{}
	// case 38:
	// 	*o = &JsonCorpse{}
	// case 39:
	// 	*o = &LootData{}
	// case 41:
	// 	*o = &SlotItemAddress{}
	// case 42:
	// 	*o = &StackSlotItemAddress{}
	// case 43:
	// 	*o = &Container{}
	// case 44:
	// 	*o = &GridItemAddress{}
	// case 45:
	// 	*o = &OwnerItself{}
	// case 46:
	// 	*o = &DestroyedItem{}
	// case 48:
	// 	*o = &AddOperation{}
	// case 49:
	// 	*o = &MagOperation{}
	// case 50:
	// 	*o = &LoadMagOperation{}
	// case 51:
	// 	*o = &UnloadMagOperation{}
	// case 52:
	// 	*o = &RemoveOperation{}
	// case 53:
	// 	*o = &ExamineOperation{}
	// case 54:
	// 	*o = &ExamineMalfunctionOperation{}
	// case 55:
	// 	*o = &ExamineMalfTypeOperation{}
	// case 56:
	// 	*o = &CheckMagazineOperation{}
	// case 57:
	// 	*o = &BindItemOperation{}
	// case 58:
	// 	*o = &UnbindItemOperation{}
	// case 59:
	// 	*o = &InsureItemsOperation{}
	// case 60:
	// 	*o = &MoveOperation{}
	// case 61:
	// 	*o = &MoveAllOperation{}
	// case 62:
	// 	*o = &SplitOperation{}
	// case 63:
	// 	*o = &MergeOperation{}
	// case 64:
	// 	*o = &TransferOperation{}
	// case 65:
	// 	*o = &SwapOperation{}
	// case 66:
	// 	*o = &ThrowOperation{}
	// case 67:
	// 	*o = &ToggleOperation{}
	// case 68:
	// 	*o = &FoldOperation{}
	// case 69:
	// 	*o = &ShotOperation{}
	// case 70:
	// 	*o = &SetupItemOperation{}
	// case 71:
	// 	*o = &ApplyOperation{}
	// case 72:
	// 	*o = &ApplyHealthOperation{}
	// case 73:
	// 	*o = &CreateMapMarkerOperation{}
	// case 74:
	// 	*o = &EditMapMarkerOperation{}
	// case 75:
	// 	*o = &DeleteMapMarkerOperation{}
	// case 76:
	// 	*o = &AddNoteOperation{}
	// case 77:
	// 	*o = &EditNoteOperation{}
	// case 78:
	// 	*o = &DeleteNoteOperation{}
	// case 79:
	// 	*o = &TagOperation{}
	// case 80:
	// 	*o = &OperateStationaryWeaponOperation{}
	// case 81:
	// 	*o = &WeaponRechamberOperation{}
	// case 82:
	// 	*o = &ObservedSyncItemsOperation{}
	// case 83:
	// 	*o = &TraderServiceAvailabilityData{}
	// case 84:
	// 	*o = &QuestAction{}
	// case 85:
	// 	*o = &QuestAccept{}
	// case 86:
	// 	*o = &QuestFinish{}
	// case 87:
	// 	*o = &QuestHandover{}
	// case 88:
	// 	*o = &SceneResourceKey{}
	// case 89:
	// 	*o = &ResourceKey{}
	// case 90:
	// 	*o = &NotesNote{}
	// case 91:
	// 	*o = &InventoryLogicMapMarker{}
	// case 92:
	// 	*o = &InventoryLogicOperationsCreateItems{}
	// case 93:
	// 	*o = &InventoryLogicOperationsPurchaseTraderServiceOperation{}
	default:
		return fmt.Errorf("invalid type to Polymorph: %d", t)
	}
	err = (*o).Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func WritePolymorph(buffer *bytes.Buffer, o Serializable) error {
	var t byte
	switch o.(type) {
	case *Quaternion:
		t = 0
	case *FoodDrinkComponent:
		t = 17
	case *MedKitComponent:
		t = 23
	case *ResourceItemComponent:
		t = 19
	default:
		return fmt.Errorf("invalid type to WritePolymorph: %T", o)
	}
	err := buffer.WriteByte(t)
	if err != nil {
		return err
	}
	return o.Serialize(buffer)
}
