package models

import "encoding/json"

type CharacterSkillsGroup struct {
	CharacterID string           `gorm:"primaryKey" json:"-"`
	Skills      []CharacterSkill `json:"-" gorm:"foreignKey:CharacterSkillsGroupID;references:CharacterID"`
	Points      int              `json:"Points"`
}

type CharacterSkill struct {
	CharacterSkillsGroupID    string   `gorm:"primaryKey" json:"-"`
	ID                        string   `gorm:"primaryKey" json:"id"`
	Type                      string   `gorm:"primaryKey" json:"-"`
	Progress                  float64  `json:"progress"`
	PointsEarnedDuringSession *float64 `json:"pointsEarnedDuringSession,omitempty"`
	LastAccess                *int32   `json:"lastAccess,omitempty"`
}

func (csg CharacterSkillsGroup) MarshalJSON() ([]byte, error) {
	type Alias CharacterSkillsGroup

	commonSkills := []CharacterSkill{}
	masterSkills := []CharacterSkill{}

	for _, skill := range csg.Skills {
		if skill.Type == "Common" {
			commonSkills = append(commonSkills, skill)
		} else if skill.Type == "Mastering" {
			masterSkills = append(masterSkills, skill)
		}
	}

	return json.Marshal(&struct {
		Common    []CharacterSkill `json:"Common"`
		Mastering []CharacterSkill `json:"Mastering"`
		*Alias
	}{
		Common:    commonSkills,
		Mastering: masterSkills,
		Alias:     (*Alias)(&csg),
	})
}
