package models

import (
	"client-server/helpers"
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	ID       int `gorm:"primaryKey" json:"uid"`
	Name     string
	Email    string    `gorm:"unique"`
	Profiles []Profile `gorm:"constraint:OnUpdate:CASCADE" json:"profiles"`
}

type BaseModel struct {
	ID string `gorm:"uniqueIndex;primaryKey" json:"-"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = helpers.GenerateUUID()
	return
}

type Profile struct {
	BaseModel
	ID                    string                     `gorm:"primaryKey" json:"_id"`
	AccountID             int                        `json:"aid" gorm:"not null;index"`
	SavageID              *string                    `json:"savage"`
	Info                  ProfileInfo                `gorm:"constraint:OnUpdate:CASCADE"`
	Customization         ProfileCustomization       `gorm:"constraint:OnUpdate:CASCADE"`
	Health                ProfileHealth              `gorm:"constraint:OnUpdate:CASCADE"`
	Inventory             ProfileInventory           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:ProfileID;references:ID"`
	Skills                ProfileSkillsGroup         `gorm:"constraint:OnUpdate:CASCADE"`
	Stats                 ProfileStats               `gorm:"constraint:OnUpdate:CASCADE"`
	Encyclopedia          ProfileEncyclopedia        `gorm:"-"`
	EncyclopediaEntries   []ProfileEncyclopediaEntry `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	TaskConditionCounters datatypes.JSON
	InsuredItems          datatypes.JSON
	Hideout               ProfileHideout        `gorm:"constraint:OnUpdate:CASCADE"`
	Bonuses               []ProfileBonus        `gorm:"constraint:OnUpdate:CASCADE"`
	Notes                 ProfileNotes          `gorm:"constraint:OnUpdate:CASCADE"`
	Quests                []ProfileQuest        `gorm:"constraint:OnUpdate:CASCADE"`
	Achievements          []ProfileAchievement  `gorm:"constraint:OnUpdate:CASCADE"`
	RagfairInfo           ProfileRagfairInfo    `gorm:"constraint:OnUpdate:CASCADE"`
	WishList              []ProfileWishListItem `gorm:"constraint:OnUpdate:CASCADE"`
	TradersInfo           ProfileTradersInfo    `gorm:"constraint:OnUpdate:CASCADE"`
	UnlockedInfo          ProfileUnlockedInfo   `gorm:"constraint:OnUpdate:CASCADE"`
	Status                ProfileStatus         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:ProfileID" json:"-"`
}

type ProfileCustomization struct {
	ProfileID string `gorm:"primaryKey" json:"-"`
	Head      string
	Body      string
	Feet      string
	Hands     string
}

type ProfileHideout struct {
	ProfileID    string               `gorm:"primaryKey" json:"-"`
	Production   datatypes.JSON       `json:"Production"`
	Areas        []ProfileHideoutArea `json:"Areas"`
	Improvements datatypes.JSON       `json:"Improvements"`
	Seed         int                  `json:"Seed"`
}

type ProfileHideoutArea struct {
	ProfileHideoutID      string         `gorm:"primaryKey" json:"-"`
	Type                  int            `gorm:"primaryKey" json:"type"`
	Level                 int            `json:"level"`
	Active                bool           `json:"active"`
	PassiveBonusesEnabled bool           `json:"passiveBonusesEnabled"`
	CompleteTime          int            `json:"completeTime"`
	Constructing          bool           `json:"constructing"`
	Slots                 datatypes.JSON `json:"slots"`
	LastRecipe            string         `json:"lastRecipe"`
}

type ProfileBonus struct {
	ProfileID  string `gorm:"primaryKey" json:"-"`
	ID         string `gorm:"primaryKey" json:"id"`
	Type       string `json:"type"`
	TemplateID string `json:"templateId"`
}

type ProfileNotes struct {
	ProfileID string         `gorm:"primaryKey" json:"-"`
	Notes     datatypes.JSON `json:"Notes"`
}

type ProfileQuest struct {
	ProfileID string `gorm:"primaryKey" json:"-"`
	QuestID   string `gorm:"primaryKey"`
}

type ProfileRagfairInfo struct {
	ProfileID       string         `gorm:"primaryKey" json:"-"`
	Rating          float64        `json:"rating"`
	IsRatingGrowing bool           `json:"isRatingGrowing"`
	Offers          datatypes.JSON `json:"offers"`
}

type ProfileWishListItem struct {
	ProfileID string `gorm:"primaryKey" json:"-"`
	ItemID    string `gorm:"primaryKey"`
}

type ProfileUnlockedInfo struct {
	ProfileID                string         `gorm:"primaryKey" json:"-"`
	UnlockedProductionRecipe datatypes.JSON `json:"unlockedProductionRecipe"`
}

type ProfileStatus struct {
	ProfileID    string  `gorm:"primaryKey" json:"profileid"`
	ProfileToken *string `json:"profiletoken"`
	Status       string  `json:"status" gorm:"default:Free"`
	ServerId     string  `json:"sid"`
	IP           string  `json:"ip"`
	Port         int     `json:"port"`
}

func (c Profile) MarshalJSON() ([]byte, error) {
	type Alias Profile
	achievements := make(ProfileAchievementsMap)
	for _, achievement := range c.Achievements {
		achievements[achievement.Name] = achievement.Value
	}

	encyclopedia := make(ProfileEncyclopedia)
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
		Achievements ProfileAchievementsMap `json:"Achievements"`
		Encyclopedia ProfileEncyclopedia    `json:"Encyclopedia"`
		TradersInfo  TradersInfoMap         `json:"TradersInfo"`
		*Alias
	}{
		Encyclopedia: encyclopedia,
		Achievements: achievements,
		TradersInfo:  tradersInfoMap,
		Alias:        (*Alias)(&c),
	})
}

func (c *Profile) UnmarshalJSON(data []byte) error {
	type Alias Profile
	aux := &struct {
		Achievements ProfileAchievementsMap `json:"Achievements"`
		TradersInfo  TradersInfoMap         `json:"TradersInfo"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	for name, value := range aux.Achievements {
		c.Achievements = append(c.Achievements, ProfileAchievement{
			ProfileID: c.ID,
			Name:      name,
			Value:     value,
		})
	}

	c.TradersInfo.Traders = make([]ProfileTraderInfo, 0, len(aux.TradersInfo))
	for traderId, traderInfo := range aux.TradersInfo {
		c.TradersInfo.Traders = append(c.TradersInfo.Traders, ProfileTraderInfo{
			ProfileID: c.TradersInfo.ProfileID,
			TraderID:  traderId,
			Unlocked:  traderInfo.Unlocked,
			Disabled:  traderInfo.Disabled,
			SalesSum:  traderInfo.SalesSum,
			Standing:  traderInfo.Standing,
		})
	}

	return nil
}
