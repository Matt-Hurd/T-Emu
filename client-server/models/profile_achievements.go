package models

import "encoding/json"

type ProfileAchievement struct {
	ProfileID string `json:"-" gorm:"primaryKey"`
	Name      string `gorm:"primaryKey" json:"name"`
	Value     int    `json:"value"`
}

type ProfileAchievementsMap map[string]int

func (a ProfileAchievementsMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]int(a))
}

func (a *ProfileAchievementsMap) UnmarshalJSON(data []byte) error {
	var temp map[string]int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*a = ProfileAchievementsMap(temp)
	return nil
}
