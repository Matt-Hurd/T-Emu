package models

import "encoding/json"

type ProfileSkillsGroup struct {
	ProfileID string         `gorm:"primaryKey" json:"-"`
	Skills    []ProfileSkill `json:"-" gorm:"foreignKey:ProfileSkillsGroupID;references:ProfileID"`
	Points    int            `json:"Points"`
}

type ProfileSkill struct {
	ProfileSkillsGroupID      string   `gorm:"primaryKey" json:"-"`
	ID                        string   `gorm:"primaryKey" json:"id"`
	Type                      string   `gorm:"primaryKey" json:"-"`
	Progress                  float64  `json:"progress"`
	PointsEarnedDuringSession *float64 `json:"pointsEarnedDuringSession,omitempty"`
	LastAccess                *int32   `json:"lastAccess,omitempty"`
}

func (csg ProfileSkillsGroup) MarshalJSON() ([]byte, error) {
	type Alias ProfileSkillsGroup

	commonSkills := []ProfileSkill{}
	masterSkills := []ProfileSkill{}

	for _, skill := range csg.Skills {
		if skill.Type == "Common" {
			commonSkills = append(commonSkills, skill)
		} else if skill.Type == "Mastering" {
			masterSkills = append(masterSkills, skill)
		}
	}

	return json.Marshal(&struct {
		Common    []ProfileSkill `json:"Common"`
		Mastering []ProfileSkill `json:"Mastering"`
		*Alias
	}{
		Common:    commonSkills,
		Mastering: masterSkills,
		Alias:     (*Alias)(&csg),
	})
}
