package models

import "encoding/json"

type ProfileHealth struct {
	ProfileID   string                 `gorm:"primaryKey" json:"-"`
	Hydration   ProfileHealthAttribute `gorm:"embedded;embeddedPrefix:hydration_"`
	Energy      ProfileHealthAttribute `gorm:"embedded;embeddedPrefix:energy_"`
	Temperature ProfileHealthAttribute `gorm:"embedded;embeddedPrefix:temperature_"`
	BodyParts   []ProfileBodyPart      `json:"-" gorm:"foreignKey:ProfileID"`
	UpdateTime  int64
	Immortal    bool `gorm:"default:false"`
}

type ProfileHealthAttribute struct {
	Current float64
	Maximum float64
}

type ProfileBodyPart struct {
	ProfileID string                 `gorm:"primaryKey" json:"-"`
	Name      string                 `gorm:"primaryKey" json:"-"`
	Health    ProfileHealthAttribute `gorm:"embedded;embeddedPrefix:health_"`
}

func (c ProfileHealth) MarshalJSON() ([]byte, error) {
	type Alias ProfileHealth
	bodyPartsMap := make(map[string]ProfilePart)
	for _, part := range c.BodyParts {
		bodyPartsMap[part.Name] = ProfilePart{
			Health: part.Health,
		}
	}
	return json.Marshal(&struct {
		Alias
		BodyParts map[string]ProfilePart `json:"BodyParts"`
	}{
		Alias:     (Alias)(c),
		BodyParts: bodyPartsMap,
	})
}

type ProfilePart struct {
	Health ProfileHealthAttribute
}
