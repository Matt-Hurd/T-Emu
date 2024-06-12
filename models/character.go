package models

import (
	"eft-private-server/helpers"
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	ID         int `gorm:"primaryKey" json:"uid"`
	Name       string
	Email      string      `gorm:"unique"`
	Characters []Character `gorm:"constraint:OnUpdate:CASCADE" json:"characters"`
}

type BaseModel struct {
	ID string `gorm:"uniqueIndex;primaryKey" json:"-"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = helpers.GenerateUUID()
	return
}

type Character struct {
	BaseModel
	ID                    string                       `gorm:"primaryKey" json:"_id"`
	AccountID             int                          `json:"aid" gorm:"not null;index"`
	SavageID              *string                      `json:"savage"`
	Info                  CharacterInfo                `gorm:"constraint:OnUpdate:CASCADE"`
	Customization         CharacterCustomization       `gorm:"constraint:OnUpdate:CASCADE"`
	Health                CharacterHealth              `gorm:"constraint:OnUpdate:CASCADE"`
	Inventory             CharacterInventory           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:CharacterID;references:ID"`
	Skills                CharacterSkillsGroup         `gorm:"constraint:OnUpdate:CASCADE"`
	Stats                 CharacterStats               `gorm:"constraint:OnUpdate:CASCADE"`
	Encyclopedia          CharacterEncyclopedia        `gorm:"-"`
	EncyclopediaEntries   []CharacterEncyclopediaEntry `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	TaskConditionCounters datatypes.JSON
	InsuredItems          datatypes.JSON
	Hideout               CharacterHideout        `gorm:"constraint:OnUpdate:CASCADE"`
	Bonuses               []CharacterBonus        `gorm:"constraint:OnUpdate:CASCADE"`
	Notes                 CharacterNotes          `gorm:"constraint:OnUpdate:CASCADE"`
	Quests                []CharacterQuest        `gorm:"constraint:OnUpdate:CASCADE"`
	Achievements          []CharacterAchievement  `gorm:"constraint:OnUpdate:CASCADE"`
	RagfairInfo           CharacterRagfairInfo    `gorm:"constraint:OnUpdate:CASCADE"`
	WishList              []CharacterWishListItem `gorm:"constraint:OnUpdate:CASCADE"`
	TradersInfo           CharacterTradersInfo    `gorm:"constraint:OnUpdate:CASCADE"`
	UnlockedInfo          CharacterUnlockedInfo   `gorm:"constraint:OnUpdate:CASCADE"`
}

type CharacterCustomization struct {
	CharacterID string `gorm:"primaryKey" json:"-"`
	Head        string
	Body        string
	Feet        string
	Hands       string
}

type CharacterHideout struct {
	CharacterID  string                 `gorm:"primaryKey" json:"-"`
	Production   datatypes.JSON         `json:"Production"`
	Areas        []CharacterHideoutArea `json:"Areas"`
	Improvements datatypes.JSON         `json:"Improvements"`
	Seed         int                    `json:"Seed"`
}

type CharacterHideoutArea struct {
	CharacterHideoutID    string         `gorm:"primaryKey" json:"-"`
	Type                  int            `json:"type"`
	Level                 int            `json:"level"`
	Active                bool           `json:"active"`
	PassiveBonusesEnabled bool           `json:"passiveBonusesEnabled"`
	CompleteTime          int            `json:"completeTime"`
	Constructing          bool           `json:"constructing"`
	Slots                 datatypes.JSON `json:"slots"`
	LastRecipe            string         `json:"lastRecipe"`
}

type CharacterBonus struct {
	CharacterID string `gorm:"primaryKey" json:"-"`
	ID          string `gorm:"primaryKey" json:"id"`
	Type        string `json:"type"`
	TemplateID  string `json:"templateId"`
}

type CharacterNotes struct {
	CharacterID string         `gorm:"primaryKey" json:"-"`
	Notes       datatypes.JSON `json:"Notes"`
}

type CharacterQuest struct {
	CharacterID string `gorm:"primaryKey" json:"-"`
	QuestID     string `gorm:"primaryKey"`
}

type CharacterRagfairInfo struct {
	CharacterID     string         `gorm:"primaryKey" json:"-"`
	Rating          float64        `json:"rating"`
	IsRatingGrowing bool           `json:"isRatingGrowing"`
	Offers          datatypes.JSON `json:"offers"`
}

type CharacterWishListItem struct {
	CharacterID string `gorm:"primaryKey" json:"-"`
	ItemID      string `gorm:"primaryKey"`
}

type CharacterUnlockedInfo struct {
	CharacterID              string         `gorm:"primaryKey" json:"-"`
	UnlockedProductionRecipe datatypes.JSON `json:"unlockedProductionRecipe"`
}

func (c Character) MarshalJSON() ([]byte, error) {
	type Alias Character
	achievements := make(CharacterAchievementsMap)
	for _, achievement := range c.Achievements {
		achievements[achievement.Name] = achievement.Value
	}

	encyclopedia := make(CharacterEncyclopedia)
	for _, entry := range c.EncyclopediaEntries {
		encyclopedia[entry.TplID] = true
	}

	blacklist := map[string]bool{
		"55d7217a4bdc2d86028b456d": true,
		"5963866286f7747bf429b572": true,
		"5963866b86f7747bfa1c4462": true,
		"602543c13fee350cd564d032": true,
	}

	// Add inventory items as false if not already in encyclopedia
	for _, item := range c.Inventory.Items {
		if _, exists := encyclopedia[item.Tpl]; !exists && !blacklist[item.Tpl] {
			encyclopedia[item.Tpl] = false
		}
	}

	fmt.Printf("c.TradersInfo: %v\n", c.TradersInfo)
	tradersInfoMap := make(TradersInfoMap)
	for _, traderInfo := range c.TradersInfo.Traders {
		tradersInfoMap[traderInfo.TraderID] = traderInfo
	}

	return json.Marshal(&struct {
		Achievements CharacterAchievementsMap `json:"Achievements"`
		Encyclopedia CharacterEncyclopedia    `json:"Encyclopedia"`
		TradersInfo  TradersInfoMap           `json:"TradersInfo"`
		*Alias
	}{
		Encyclopedia: encyclopedia,
		Achievements: achievements,
		TradersInfo:  tradersInfoMap,
		Alias:        (*Alias)(&c),
	})
}

func (c *Character) UnmarshalJSON(data []byte) error {
	type Alias Character
	aux := &struct {
		Achievements CharacterAchievementsMap `json:"Achievements"`
		TradersInfo  TradersInfoMap           `json:"TradersInfo"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	for name, value := range aux.Achievements {
		c.Achievements = append(c.Achievements, CharacterAchievement{
			CharacterID: c.ID,
			Name:        name,
			Value:       value,
		})
	}

	c.TradersInfo.Traders = make([]CharacterTraderInfo, 0, len(aux.TradersInfo))
	for traderId, traderInfo := range aux.TradersInfo {
		c.TradersInfo.Traders = append(c.TradersInfo.Traders, CharacterTraderInfo{
			CharacterID: c.TradersInfo.CharacterID,
			TraderID:    traderId,
			Unlocked:    traderInfo.Unlocked,
			Disabled:    traderInfo.Disabled,
			SalesSum:    traderInfo.SalesSum,
			Standing:    traderInfo.Standing,
		})
	}

	return nil
}
