package descriptors

import (
	"bytes"
	"game-server/helpers"
)

// GClass1512
type FoodDrinkComponentDescriptor struct {
	HpPercent float32
}

func (foodDrinkComponent *FoodDrinkComponentDescriptor) Serialize(buffer *bytes.Buffer) error {
	err := helpers.WriteFloat32(buffer, foodDrinkComponent.HpPercent)
	if err != nil {
		return err
	}
	return nil
}

func (foodDrinkComponent *FoodDrinkComponentDescriptor) Deserialize(buffer *bytes.Buffer) error {
	err := helpers.ReadFloat32(buffer, &foodDrinkComponent.HpPercent)
	if err != nil {
		return err
	}
	return nil
}
