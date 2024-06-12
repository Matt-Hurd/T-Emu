package models

type ProfileTradersInfo struct {
	ProfileID string              `gorm:"primaryKey" json:"-"`
	Traders   []ProfileTraderInfo `json:"-" gorm:"foreignKey:ProfileID"`
}

type ProfileTraderInfo struct {
	ProfileID string  `json:"-" gorm:"primaryKey"`
	TraderID  string  `json:"-" gorm:"primaryKey"`
	Unlocked  bool    `json:"unlocked"`
	Disabled  bool    `json:"disabled"`
	SalesSum  int     `json:"salesSum"`
	Standing  float64 `json:"standing"`
}

type TradersInfoMap map[string]ProfileTraderInfo
