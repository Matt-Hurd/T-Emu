package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringArray []string

// Scan implements the Scanner interface for deserialization
func (a *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, a)
}

// Value implements the Valuer interface for serialization
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type CharacterStats struct {
	CharacterID string `gorm:"primaryKey" json:"-"`
	Eft         CharacterEftStats
}

type CharacterEftStats struct {
	CharacterStatsID       string `gorm:"primaryKey" json:"-"`
	SessionCounters        CharacterCounter
	OverallCounters        CharacterCounter
	SessionExperienceMult  float64
	ExperienceBonusMult    float64
	TotalSessionExperience int
	LastSessionDate        int64
	Aggressor              *string
	DroppedItems           StringArray
	FoundInRaidItems       StringArray
	Victims                StringArray
	CarriedQuestItems      StringArray
	DamageHistory          CharacterDamageHistory
	LastPlayerState        *string
	TotalInGameTime        int
	SurvivorClass          string
}

type CharacterCounter struct {
	CharacterEftStatsID string                 `gorm:"primaryKey" json:"-"`
	Items               []CharacterCounterItem `json:"Items"`
}

type CharacterCounterItem struct {
	CharacterCounterID string `gorm:"primaryKey" json:"-"`
	Key                StringArray
	Value              int
}

type CharacterDamageHistory struct {
	BaseModel
	CharacterEftStatsID string `gorm:"primaryKey" json:"-"`
	LethalDamagePart    string `gorm:"default:Head"`
	LethalDamage        *CharacterBodyPartDamage
	BodyParts           []CharacterBodyPartDamage
}

type CharacterBodyPartDamage struct {
	ID                       uint `gorm:"primaryKey"`
	CharacterDamageHistoryID uint
	BodyPart                 string
	Amount                   int
	Type                     string
	SourceID                 *string
	OverDamageFrom           *string
	Blunt                    bool
	ImpactsCount             int
}

func (cdh CharacterDamageHistory) MarshalJSON() ([]byte, error) {
	type Alias CharacterDamageHistory
	bodyParts := []string{"Head", "Chest", "Stomach", "LeftArm", "RightArm", "LeftLeg", "RightLeg"}
	bodyPartsMap := make(map[string][]CharacterBodyPartDamage)
	for _, part := range bodyParts {
		bodyPartsMap[part] = []CharacterBodyPartDamage{}
	}

	for _, partDamage := range cdh.BodyParts {
		bodyPartsMap[partDamage.BodyPart] = append(bodyPartsMap[partDamage.BodyPart], partDamage)
	}

	return json.Marshal(&struct {
		BodyParts map[string][]CharacterBodyPartDamage
		*Alias
	}{
		BodyParts: bodyPartsMap,
		Alias:     (*Alias)(&cdh),
	})
}
