package models

import "encoding/json"

type CharacterTradersInfo struct {
	CharacterID string                `gorm:"primaryKey" json:"-"`
	Traders     []CharacterTraderInfo `json:"-" gorm:"foreignKey:CharacterID;references:CharacterID"`
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

func (cti CharacterTradersInfo) MarshalJSON() ([]byte, error) {
	type Alias CharacterTradersInfo

	tradersInfoMap := make(TradersInfoMap)
	for _, traderInfo := range cti.Traders {
		tradersInfoMap[traderInfo.TraderID] = traderInfo
	}

	return json.Marshal(&struct {
		TradersInfo TradersInfoMap `json:"TradersInfo"`
		*Alias
	}{
		TradersInfo: tradersInfoMap,
		Alias:       (*Alias)(&cti),
	})
}

func (cti *CharacterTradersInfo) UnmarshalJSON(data []byte) error {
	type Alias CharacterTradersInfo
	aux := &struct {
		TradersInfo TradersInfoMap `json:"TradersInfo"`
		*Alias
	}{
		Alias: (*Alias)(cti),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	for traderId, traderInfo := range aux.TradersInfo {
		cti.Traders = append(cti.Traders, CharacterTraderInfo{
			CharacterID: cti.CharacterID,
			TraderID:    traderId,
			Unlocked:    traderInfo.Unlocked,
			Disabled:    traderInfo.Disabled,
			SalesSum:    traderInfo.SalesSum,
			Standing:    traderInfo.Standing,
		})
	}

	return nil
}
