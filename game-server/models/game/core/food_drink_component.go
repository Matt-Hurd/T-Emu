package core

import (
	"bytes"
	"game-server/helpers"
)

// GClass1512
type FoodDrinkComponent struct {
	HpPercent float32
}

func (foodDrinkComponent *FoodDrinkComponent) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, foodDrinkComponent.HpPercent)
	if err != nil {
		return err
	}
	return nil
}

func (foodDrinkComponent *FoodDrinkComponent) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &foodDrinkComponent.HpPercent)
	if err != nil {
		return err
	}
	return nil
}
