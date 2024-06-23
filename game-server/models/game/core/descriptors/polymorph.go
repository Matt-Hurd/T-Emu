package descriptors

import (
	"bytes"
	"fmt"
	"game-server/helpers"
	"game-server/models/game/core"
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
		*o = &core.Quaternion{}
	// case 1:
	// 	*o = &ClassTransformSync{}
	case 2:
		*o = &core.Vector3{}
	case 3:
		*o = &LocationInGrid{}
	// case 4:
	// 	*o = &WeightedLootPointSpawnPosition{}
	case 5:
		*o = &InventoryDescriptor{}
	case 6:
		*o = &FastAccessDescriptor{}
	case 7:
		*o = &DiscardLimitsDescriptor{}
	case 8:
		*o = &SlotDescriptor{}
	// case 9:
	// 	*o = &ShellTemplateDescriptor{}
	case 10:
		*o = &MalfunctionDescriptor{}
	// case 11:
	// 	*o = &ItemInGridDescriptor{}
	// case 12:
	// 	*o = &GridDescriptor{}
	case 13:
		*o = &StackSlotDescriptor{}
	// case 14:
	// 	*o = &NestedItemDescriptor{}
	case 15:
		*o = &ItemDescriptor{}
	case 17:
		*o = &FoodDrinkComponentDescriptor{}
	// case 18:
	// 	*o = &PoisonComponentDescriptor{}
	case 19:
		*o = &ResourceItemComponentDescriptor{}
	// case 20:
	// 	*o = &LightComponentDescriptor{}
	// case 21:
	// 	*o = &LockableComponentDescriptor{}
	// case 22:
	// 	*o = &MapComponentDescriptor{}
	case 23:
		*o = &MedKitComponentDescriptor{}
	// case 24:
	// 	*o = &RepairableComponentDescriptor{}
	// case 25:
	// 	*o = &SightComponentDescriptor{}
	// case 26:
	// 	*o = &TogglableComponentDescriptor{}
	// case 27:
	// 	*o = &FaceShieldComponentDescriptor{}
	// case 28:
	// 	*o = &FoldableComponentDescriptor{}
	// case 29:
	// 	*o = &FireModeComponentDescriptor{}
	// case 30:
	// 	*o = &DogTagComponentDescriptor{}
	// case 31:
	// 	*o = &TagComponentDescriptor{}
	// case 32:
	// 	*o = &KeyComponentDescriptor{}
	// case 33:
	// 	*o = &RepairKitComponentDescriptor{}
	// case 34:
	// 	*o = &RepairEnhancementComponentDescriptor{}
	// case 35:
	// 	*o = &RecodableComponentDescriptor{}
	// case 36:
	// 	*o = &CultistAmuletComponentDescriptor{}
	case 37:
		*o = &LootDescriptor{}
	// case 38:
	// 	*o = &JsonCorpseDescriptor{}
	// case 39:
	// 	*o = &LootDataDescriptor{}
	// case 41:
	// 	*o = &SlotItemAddressDescriptor{}
	// case 42:
	// 	*o = &StackSlotItemAddressDescriptor{}
	// case 43:
	// 	*o = &ContainerDescriptor{}
	// case 44:
	// 	*o = &GridItemAddressDescriptor{}
	// case 45:
	// 	*o = &OwnerItselfDescriptor{}
	// case 46:
	// 	*o = &DestroyedItem{}
	// case 48:
	// 	*o = &AddOperationDescriptor{}
	// case 49:
	// 	*o = &MagOperationDescriptor{}
	// case 50:
	// 	*o = &LoadMagOperationDescriptor{}
	// case 51:
	// 	*o = &UnloadMagOperationDescriptor{}
	// case 52:
	// 	*o = &RemoveOperationDescriptor{}
	// case 53:
	// 	*o = &ExamineOperationDescriptor{}
	// case 54:
	// 	*o = &ExamineMalfunctionOperationDescriptor{}
	// case 55:
	// 	*o = &ExamineMalfTypeOperationDescriptor{}
	// case 56:
	// 	*o = &CheckMagazineOperationDescriptor{}
	// case 57:
	// 	*o = &BindItemOperationDescriptor{}
	// case 58:
	// 	*o = &UnbindItemOperationDescriptor{}
	// case 59:
	// 	*o = &InsureItemsOperationDescriptor{}
	// case 60:
	// 	*o = &MoveOperationDescriptor{}
	// case 61:
	// 	*o = &MoveAllOperationDescriptor{}
	// case 62:
	// 	*o = &SplitOperationDescriptor{}
	// case 63:
	// 	*o = &MergeOperationDescriptor{}
	// case 64:
	// 	*o = &TransferOperationDescriptor{}
	// case 65:
	// 	*o = &SwapOperationDescriptor{}
	// case 66:
	// 	*o = &ThrowOperationDescriptor{}
	// case 67:
	// 	*o = &ToggleOperationDescriptor{}
	// case 68:
	// 	*o = &FoldOperationDescriptor{}
	// case 69:
	// 	*o = &ShotOperationDescriptor{}
	// case 70:
	// 	*o = &SetupItemOperationDescriptor{}
	// case 71:
	// 	*o = &ApplyOperationDescriptor{}
	// case 72:
	// 	*o = &ApplyHealthOperationDescriptor{}
	// case 73:
	// 	*o = &CreateMapMarkerOperationDescriptor{}
	// case 74:
	// 	*o = &EditMapMarkerOperationDescriptor{}
	// case 75:
	// 	*o = &DeleteMapMarkerOperationDescriptor{}
	// case 76:
	// 	*o = &AddNoteOperationDescriptor{}
	// case 77:
	// 	*o = &EditNoteOperationDescriptor{}
	// case 78:
	// 	*o = &DeleteNoteOperationDescriptor{}
	// case 79:
	// 	*o = &TagOperationDescriptor{}
	// case 80:
	// 	*o = &OperateStationaryWeaponOperationDescriptor{}
	// case 81:
	// 	*o = &WeaponRechamberOperationDescriptor{}
	// case 82:
	// 	*o = &ObservedSyncItemsOperationDescriptor{}
	// case 83:
	// 	*o = &TraderServiceAvailabilityData{}
	// case 84:
	// 	*o = &QuestActionDescriptor{}
	// case 85:
	// 	*o = &QuestAcceptDescriptor{}
	// case 86:
	// 	*o = &QuestFinishDescriptor{}
	// case 87:
	// 	*o = &QuestHandoverDescriptor{}
	// case 88:
	// 	*o = &SceneResourceKey{}
	// case 89:
	// 	*o = &ResourceKey{}
	// case 90:
	// 	*o = &NotesNote{}
	// case 91:
	// 	*o = &InventoryLogicMapMarker{}
	// case 92:
	// 	*o = &InventoryLogicOperationsCreateItemsDescriptor{}
	// case 93:
	// 	*o = &InventoryLogicOperationsPurchaseTraderServiceOperationDescriptor{}
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
	case *core.Quaternion:
		t = 0
	case *core.Vector3:
		t = 2
	case *LocationInGrid:
		t = 3
	case *InventoryDescriptor:
		t = 5
	case *FastAccessDescriptor:
		t = 6
	case *DiscardLimitsDescriptor:
		t = 7
	case *SlotDescriptor:
		t = 8
	case *MalfunctionDescriptor:
		t = 10
	case *StackSlotDescriptor:
		t = 13
	case *ItemDescriptor:
		t = 15
	case *FoodDrinkComponentDescriptor:
		t = 17
	case *ResourceItemComponentDescriptor:
		t = 19
	case *MedKitComponentDescriptor:
		t = 23
	case *LootDescriptor:
		t = 37
	default:
		return fmt.Errorf("invalid type to WritePolymorph: %T", o)
	}
	err := buffer.WriteByte(t)
	if err != nil {
		return err
	}
	return o.Serialize(buffer)
}
