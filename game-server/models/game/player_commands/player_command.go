package playercommands

import (
	"bytes"
	"fmt"
	"game-server/models/game/enums"
)

type PlayerCommand struct {
	Type enums.CommandMessageType
	Data PlayerCommandMsg
}

type PlayerCommandMsg interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

func (command *PlayerCommand) Serialize(buffer *bytes.Buffer) error {
	buffer.WriteByte(byte(command.Type))
	err := command.Data.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (command *PlayerCommand) Deserialize(buffer *bytes.Buffer) error {
	val, err := buffer.ReadByte()
	command.Type = enums.CommandMessageType(val)
	if err != nil {
		return err
	}
	switch command.Type {

	case enums.MsgDeath:
		command.Data = &PlayerCommandDeath{}
	case enums.MsgPhrase:
		command.Data = &PlayerCommandPhrase{}
	case enums.MsgSkillParameters:
		command.Data = &PlayerCommandSkillParameters{}
	case enums.MsgTemperature:
		command.Data = &PlayerCommandTemperature{}
	case enums.MsgHealthStatus:
		command.Data = &PlayerCommandHealthStatus{}
	case enums.MsgMedEffectStatus:
		command.Data = &PlayerCommandMedEffectStatus{}
	case enums.MsgMedEffectResource:
		command.Data = &PlayerCommandMedEffectResource{}
	case enums.MsgPhysicalParameters:
		command.Data = &PlayerCommandPhysicalParameters{}
	case enums.MsgVoIP:
		command.Data = &PlayerCommandVoIP{}
	case enums.MsgTakeDamage:
		command.Data = &PlayerCommandTakeDamage{}
	case enums.MsgHeadDeviceInfo:
		command.Data = &PlayerCommandHeadDeviceInfo{}
	case enums.MsgEffectOnPlayer:
		command.Data = &PlayerCommandEffectOnPlayer{}
	case enums.MsgSetAnimationLayerWeight:
		command.Data = &PlayerCommandSetAnimationLayerWeight{}
	case enums.MsgSetHands:
		command.Data = &PlayerCommandSetHands{}
	case enums.MsgShotFirearm:
		command.Data = &PlayerCommandShotFirearm{}
	case enums.MsgDryShotFirearm:
		command.Data = &PlayerCommandDryShotFirearm{}
	case enums.MsgFlareShotFirearm:
		command.Data = &PlayerCommandFlareShotFirearm{}
	case enums.MsgSwingGrenade:
		command.Data = &PlayerCommandSwingGrenade{}
	case enums.MsgThrowGrenade:
		command.Data = &PlayerCommandThrowGrenade{}
	case enums.MsgGesture:
		command.Data = &PlayerCommandGesture{}
	case enums.MsgArmorChangedInfo:
		command.Data = &PlayerCommandArmorChangedInfo{}
	case enums.MsgArmorDurabilityChanged:
		command.Data = &PlayerCommandArmorDurabilityChanged{}
	case enums.MsgEquipChanged:
		command.Data = &PlayerCommandEquipChanged{}
	case enums.MsgExamineWeapon:
		command.Data = &PlayerCommandExamineWeapon{}
	case enums.MsgChangeFireMode:
		command.Data = &PlayerCommandChangeFireMode{}
	case enums.MsgTacticalModeToggle:
		command.Data = &PlayerCommandTacticalModeToggle{}
	case enums.MsgScopeModeToggle:
		command.Data = &PlayerCommandScopeModeToggle{}
	case enums.MsgModChanged:
		command.Data = &PlayerCommandModChanged{}
	case enums.MsgSingleBarrelReloadStart:
		command.Data = &PlayerCommandSingleBarrelReloadStart{}
	case enums.MsgSetChamberState:
		command.Data = &PlayerCommandSetChamberState{}
	case enums.MsgMalfunction:
		command.Data = &PlayerCommandMalfunction{}
	case enums.MsgRepairMalfunction:
		command.Data = &PlayerCommandRepairMalfunction{}
	case enums.MsgReloadExternalMagazine:
		command.Data = &PlayerCommandReloadExternalMagazine{}
	case enums.MsgSetExternalMagazineState:
		command.Data = &PlayerCommandSetExternalMagazineState{}
	case enums.MsgReloadInternalMagazine:
		command.Data = &PlayerCommandReloadInternalMagazine{}
	case enums.MsgReloadCylinderMagazine:
		command.Data = &PlayerCommandReloadCylinderMagazine{}
	case enums.MsgAbortReloadCylinderMagazine:
		command.Data = &PlayerCommandAbortReloadCylinderMagazine{}
	case enums.MsgSyncCylinderMagazine:
		command.Data = &PlayerCommandSyncCylinderMagazine{}
	case enums.MsgRollCylinder:
		command.Data = &PlayerCommandRollCylinder{}
	case enums.MsgCylinderCamoraIndex:
		command.Data = &PlayerCommandCylinderCamoraIndex{}
	case enums.MsgRechamber:
		command.Data = &PlayerCommandRechamber{}
	case enums.MsgLoadAmmoToChamber:
		command.Data = &PlayerCommandLoadAmmoToChamber{}
	case enums.MsgDischargeChamber:
		command.Data = &PlayerCommandDischargeChamber{}
	case enums.MsgLoadAmmoToCamora:
		command.Data = &PlayerCommandLoadAmmoToCamora{}
	case enums.MsgDischargeAmmoFromCamora:
		command.Data = &PlayerCommandDischargeAmmoFromCamora{}
	case enums.MsgToggleUnderbarrel:
		command.Data = &PlayerCommandToggleUnderbarrel{}
	case enums.MsgUnderbarrelShot:
		command.Data = &PlayerCommandUnderbarrelShot{}
	case enums.MsgUnderbarrelReload:
		command.Data = &PlayerCommandUnderbarrelReload{}
	case enums.MsgUnderbarrelRangeValue:
		command.Data = &PlayerCommandUnderbarrelRangeValue{}
	case enums.MsgLoadAmmoToUnderbarrel:
		command.Data = &PlayerCommandLoadAmmoToUnderbarrel{}
	case enums.MsgDischargeAmmoFromUnderbarrel:
		command.Data = &PlayerCommandDischargeAmmoFromUnderbarrel{}
	case enums.MsgBoltActionReloadAfterFire:
		command.Data = &PlayerCommandBoltActionReloadAfterFire{}
	case enums.MsgInteract:
		command.Data = &PlayerCommandInteract{}
	case enums.MsgDoorUnlockInteraction:
		command.Data = &PlayerCommandDoorUnlockInteraction{}
	case enums.MsgDoorBreachInteraction:
		command.Data = &PlayerCommandDoorBreachInteraction{}
	case enums.MsgBtrGoInInteraction:
		command.Data = &PlayerCommandBtrGoInInteraction{}
	case enums.MsgBtrGoOutInteraction:
		command.Data = &PlayerCommandBtrGoOutInteraction{}
	case enums.MsgInventoryOpenStatus:
		command.Data = &PlayerCommandInventoryOpenStatus{}
	case enums.MsgAiming:
		command.Data = &PlayerCommandAiming{}
	case enums.MsgCompassState:
		command.Data = &PlayerCommandCompassState{}
	case enums.MsgMeleeAttack:
		command.Data = &PlayerCommandMeleeAttack{}
	case enums.MsgBreakMeleeCombo:
		command.Data = &PlayerCommandBreakMeleeCombo{}
	case enums.MsgIdleStateSync:
		command.Data = &PlayerCommandIdleStateSync{}
	case enums.MsgRadioTransmitterStatus:
		command.Data = &PlayerCommandRadioTransmitterStatus{}
	case enums.MsgInventoryOperation:
		command.Data = &PlayerCommandInventoryOperation{}
	case enums.MsgSetStationaryWeapon:
		command.Data = &PlayerCommandSetStationaryWeapon{}
	case enums.MsgSyncStationaryMagazine:
		command.Data = &PlayerCommandSyncStationaryMagazine{}
	case enums.MsgInsertMagazine:
		command.Data = &PlayerCommandInsertMagazine{}
	case enums.MsgPullOutMagazine:
		command.Data = &PlayerCommandPullOutMagazine{}
	case enums.MsgMagAndChamberState:
		command.Data = &PlayerCommandMagAndChamberState{}
	case enums.MsgThrowItemAsLoot:
		command.Data = &PlayerCommandThrowItemAsLoot{}
	case enums.MsgReloadInternalMagWithOpenBolt:
		command.Data = &PlayerCommandReloadInternalMagWithOpenBolt{}
	case enums.MsgSetAbortReloadInternalMagWithOpenBolt:
		command.Data = &PlayerCommandSetAbortReloadInternalMagWithOpenBolt{}
	case enums.MsgSetFinishReloadInternalMagWithOpenBolt:
		command.Data = &PlayerCommandSetFinishReloadInternalMagWithOpenBolt{}
	case enums.MsgReloadMultiBarrelWeapon:
		command.Data = &PlayerCommandReloadMultiBarrelWeapon{}
	case enums.MsgPickup:
		command.Data = &PlayerCommandPickup{}
	case enums.MsgStartSearchContent:
		command.Data = &PlayerCommandStartSearchContent{}
	case enums.MsgStopSearchContent:
		command.Data = &PlayerCommandStopSearchContent{}
	case enums.MsgSetLeftStance:
		command.Data = &PlayerCommandSetLeftStance{}
	case enums.MsgSetVoiceMuffledStatus:
		command.Data = &PlayerCommandSetVoiceMuffledStatus{}
	case enums.MsgSetUnderRoofStatus:
		command.Data = &PlayerCommandSetUnderRoofStatus{}
	case enums.MsgVaulting:
		command.Data = &PlayerCommandVaulting{}
	default:
		return fmt.Errorf("unknown command type: %d", command.Type)
	}
	err = command.Data.Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}
