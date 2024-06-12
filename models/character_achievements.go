package models

import "encoding/json"

type CharacterAchievement struct {
	CharacterID string `json:"-" gorm:"primaryKey"`
	Name        string `gorm:"primaryKey" json:"name"`
	Value       int    `json:"value"`
}

type CharacterAchievementsMap map[string]int

func (a CharacterAchievementsMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]int(a))
}

func (a *CharacterAchievementsMap) UnmarshalJSON(data []byte) error {
	var temp map[string]int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*a = CharacterAchievementsMap(temp)
	return nil
}
