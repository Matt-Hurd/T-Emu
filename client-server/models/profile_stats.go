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

type ProfileStats struct {
	ProfileID string `gorm:"primaryKey" json:"-"`
	Eft       ProfileEftStats
}

type ProfileEftStats struct {
	ProfileStatsID         string `gorm:"primaryKey" json:"-"`
	SessionCounters        ProfileCounter
	OverallCounters        ProfileCounter
	SessionExperienceMult  float64
	ExperienceBonusMult    float64
	TotalSessionExperience int
	LastSessionDate        int64
	Aggressor              *string
	DroppedItems           StringArray
	FoundInRaidItems       StringArray
	Victims                StringArray
	CarriedQuestItems      StringArray
	DamageHistory          ProfileDamageHistory
	LastPlayerState        *string
	TotalInGameTime        int
	SurvivorClass          string
}

type ProfileCounter struct {
	ProfileEftStatsID string               `gorm:"primaryKey" json:"-"`
	Items             []ProfileCounterItem `json:"Items"`
}

type ProfileCounterItem struct {
	ProfileCounterID string `gorm:"primaryKey" json:"-"`
	Key              StringArray
	Value            int
}

type ProfileDamageHistory struct {
	BaseModel
	ProfileEftStatsID string `gorm:"primaryKey" json:"-"`
	LethalDamagePart  string `gorm:"default:Head"`
	LethalDamage      *ProfileBodyPartDamage
	BodyParts         []ProfileBodyPartDamage
}

type ProfileBodyPartDamage struct {
	ID                     uint `gorm:"primaryKey"`
	ProfileDamageHistoryID uint
	BodyPart               string
	Amount                 int
	Type                   string
	SourceID               *string
	OverDamageFrom         *string
	Blunt                  bool
	ImpactsCount           int
}

func (cdh ProfileDamageHistory) MarshalJSON() ([]byte, error) {
	type Alias ProfileDamageHistory
	bodyParts := []string{"Head", "Chest", "Stomach", "LeftArm", "RightArm", "LeftLeg", "RightLeg"}
	bodyPartsMap := make(map[string][]ProfileBodyPartDamage)
	for _, part := range bodyParts {
		bodyPartsMap[part] = []ProfileBodyPartDamage{}
	}

	for _, partDamage := range cdh.BodyParts {
		bodyPartsMap[partDamage.BodyPart] = append(bodyPartsMap[partDamage.BodyPart], partDamage)
	}

	return json.Marshal(&struct {
		BodyParts map[string][]ProfileBodyPartDamage
		*Alias
	}{
		BodyParts: bodyPartsMap,
		Alias:     (*Alias)(&cdh),
	})
}
