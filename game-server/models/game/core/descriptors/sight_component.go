package descriptors

import (
	"bytes"
	"game-server/helpers"
)

type SightComponentDescriptor struct {
	SelectedSightScope        int32
	ScopeSelectedModes        []int32
	ScopesSelectedCalibPoints []int32
}

func (sightComponent *SightComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, sightComponent.SelectedSightScope); err != nil {
		return err
	}
	if err := helpers.WriteInt32(buffer, int32(len(sightComponent.ScopeSelectedModes))); err != nil {
		return err
	}
	for _, v := range sightComponent.ScopeSelectedModes {
		if err := helpers.WriteInt32(buffer, v); err != nil {
			return err
		}
	}
	if err := helpers.WriteInt32(buffer, int32(len(sightComponent.ScopesSelectedCalibPoints))); err != nil {
		return err
	}
	for _, v := range sightComponent.ScopesSelectedCalibPoints {
		if err := helpers.WriteInt32(buffer, v); err != nil {
			return err
		}
	}
	return nil
}

func (sightComponent *SightComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadInt32(buffer, &sightComponent.SelectedSightScope); err != nil {
		return err
	}
	var ScopeSelectedModesLength int32
	if err := helpers.ReadInt32(buffer, &ScopeSelectedModesLength); err != nil {
		return err
	}
	sightComponent.ScopeSelectedModes = make([]int32, ScopeSelectedModesLength)
	for i := 0; i < int(ScopeSelectedModesLength); i++ {
		if err := helpers.ReadInt32(buffer, &sightComponent.ScopeSelectedModes[i]); err != nil {
			return err
		}
	}
	var ScopesSelectedCalibPointsLength int32
	if err := helpers.ReadInt32(buffer, &ScopesSelectedCalibPointsLength); err != nil {
		return err
	}
	sightComponent.ScopesSelectedCalibPoints = make([]int32, ScopesSelectedCalibPointsLength)
	for i := 0; i < int(ScopesSelectedCalibPointsLength); i++ {
		if err := helpers.ReadInt32(buffer, &sightComponent.ScopesSelectedCalibPoints[i]); err != nil {
			return err
		}
	}
	return nil
}
