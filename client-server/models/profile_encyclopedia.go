package models

import "encoding/json"

type ProfileEncyclopediaEntry struct {
	ProfileID string `gorm:"primaryKey"`
	TplID     string `gorm:"primaryKey"`
}

type ProfileEncyclopedia map[string]bool

func (e ProfileEncyclopedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]bool(e))
}

func (e *ProfileEncyclopedia) UnmarshalJSON(data []byte) error {
	var temp map[string]bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*e = ProfileEncyclopedia(temp)
	return nil
}
