package models

import "encoding/json"

type CharacterEncyclopediaEntry struct {
	CharacterID string `gorm:"primaryKey"`
	TplID       string `gorm:"primaryKey"`
}

type CharacterEncyclopedia map[string]bool

func (e CharacterEncyclopedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]bool(e))
}

func (e *CharacterEncyclopedia) UnmarshalJSON(data []byte) error {
	var temp map[string]bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*e = CharacterEncyclopedia(temp)
	return nil
}
