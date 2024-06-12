package models

import "encoding/json"

type CharacterHealth struct {
	CharacterID string                   `gorm:"primaryKey" json:"-"`
	Hydration   CharacterHealthAttribute `gorm:"embedded;embeddedPrefix:hydration_"`
	Energy      CharacterHealthAttribute `gorm:"embedded;embeddedPrefix:energy_"`
	Temperature CharacterHealthAttribute `gorm:"embedded;embeddedPrefix:temperature_"`
	BodyParts   []CharacterBodyPart      `json:"-" gorm:"foreignKey:CharacterID"`
	UpdateTime  int64
	Immortal    bool `gorm:"default:false"`
}

type CharacterHealthAttribute struct {
	Current float64
	Maximum float64
}

type CharacterBodyPart struct {
	CharacterID string                   `gorm:"primaryKey" json:"-"`
	Name        string                   `gorm:"primaryKey" json:"-"`
	Health      CharacterHealthAttribute `gorm:"embedded;embeddedPrefix:health_"`
}

func (c CharacterHealth) MarshalJSON() ([]byte, error) {
	type Alias CharacterHealth
	bodyPartsMap := make(map[string]CharacterPart)
	for _, part := range c.BodyParts {
		bodyPartsMap[part.Name] = CharacterPart{
			Health: part.Health,
		}
	}
	return json.Marshal(&struct {
		Alias
		BodyParts map[string]CharacterPart `json:"BodyParts"`
	}{
		Alias:     (Alias)(c),
		BodyParts: bodyPartsMap,
	})
}

type CharacterPart struct {
	Health CharacterHealthAttribute
}
