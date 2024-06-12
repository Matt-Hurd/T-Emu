package models

type CharacterTradersInfo struct {
	CharacterID string                `gorm:"primaryKey" json:"-"`
	Traders     []CharacterTraderInfo `json:"-" gorm:"foreignKey:CharacterID"`
}

type CharacterTraderInfo struct {
	CharacterID string  `json:"-" gorm:"primaryKey"`
	TraderID    string  `json:"-" gorm:"primaryKey"`
	Unlocked    bool    `json:"unlocked"`
	Disabled    bool    `json:"disabled"`
	SalesSum    int     `json:"salesSum"`
	Standing    float64 `json:"standing"`
}

type TradersInfoMap map[string]CharacterTraderInfo
