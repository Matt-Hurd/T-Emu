package models

import (
	"time"

	"gorm.io/datatypes"
)

type Account struct {
	ID         string `gorm:"primaryKey" json:"uid"`
	Name       string
	Email      string      `gorm:"unique"`
	Characters []Character `json:"characters"`
}

type Character struct {
	ID                    string                 `gorm:"primaryKey" json:"_id"`
	AccountID             string                 `json:"aid" gorm:"not null;index"`
	Account               Account                `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	SavageID              *string                `json:"savage" gorm:"index"`
	Savage                *Character             `gorm:"foreignKey:SavageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Info                  CharacterInfo          `json:"Info"`
	Customization         CharacterCustomization `json:"Customization"`
	Health                CharacterHealth        `json:"Health"`
	Inventory             CharacterInventory     `json:"Inventory"`
	Skills                CharacterSkills        `json:"Skills"`
	Stats                 CharacterStats         `json:"Stats"`
	Encyclopedia          datatypes.JSON         `json:"Encyclopedia"`
	TaskConditionCounters datatypes.JSON         `json:"TaskConditionCounters"`
	Hideout               CharacterHideout       `json:"Hideout"`
	Bonuses               []CharacterBonus       `json:"Bonuses"`
	Notes                 CharacterNotes         `json:"Notes"`
	Quests                []datatypes.JSON       `json:"Quests"`
	Achievements          CharacterAchievements  `json:"Achievements"`
	RagfairInfo           CharacterRagfairInfo   `json:"RagfairInfo"`
	WishList              []datatypes.JSON       `json:"WishList"`
	TradersInfo           CharacterTradersInfo   `json:"TradersInfo"`
	UnlockedInfo          CharacterUnlockedInfo  `json:"UnlockedInfo"`
}

type CharacterInfo struct {
	Nickname                string                `json:"Nickname"`
	LowerNickname           string                `json:"LowerNickname"`
	Side                    string                `json:"Side"`
	Voice                   string                `json:"Voice"`
	Level                   int                   `gorm:"default:1" json:"Level"`
	Experience              int                   `gorm:"default:0" json:"Experience"`
	RegistrationDate        time.Time             `gorm:"default:CURRENT_TIMESTAMP" json:"RegistrationDate"`
	GameVersion             string                `json:"GameVersion"`
	AccountType             int                   `gorm:"default:0" json:"AccountType"`
	MemberCategory          int                   `gorm:"default:2" json:"MemberCategory"`
	LockedMoveCommands      bool                  `json:"LockedMoveCommands"`
	SavageLockTime          time.Time             `gorm:"default:0" json:"SavageLockTime"`
	LastTimePlayedAsSavage  time.Time             `gorm:"default:0" json:"LastTimePlayedAsSavage"`
	Settings                CharacterInfoSettings `json:"Settings"`
	NicknameChangeDate      time.Time             `gorm:"default:0" json:"NicknameChangeDate"`
	NeedWipeOptions         []string              `json:"NeedWipeOptions"`
	LastCompletedWipe       *int                  `json:"LastCompletedWipe"`
	LastCompletedEvent      *int                  `json:"LastCompletedEvent"`
	BannedState             bool                  `json:"BannedState"`
	BannedUntil             time.Time             `json:"BannedUntil"`
	IsStreamerModeAvailable bool                  `gorm:"default:false" json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool                  `gorm:"default:false" json:"SquadInviteRestriction"`
	HasCoopExtension        bool                  `gorm:"default:false" json:"HasCoopExtension"`
	Bans                    []time.Time           `json:"Bans"`
}

type CharacterInfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int     `json:"Experience"`
	StandingForKill float64 `json:"StandingForKill"`
	AggressorBonus  float64 `json:"AggressorBonus"`
}

type CharacterCustomization struct {
	Head  string `json:"Head"`
	Body  string `json:"Body"`
	Feet  string `json:"Feet"`
	Hands string `json:"Hands"`
}

type CharacterHealth struct {
	Hydration   CharacterHealthAttribute `json:"Hydration" gorm:"embedded"`
	Energy      CharacterHealthAttribute `json:"Energy" gorm:"embedded"`
	Temperature CharacterHealthAttribute `json:"Temperature" gorm:"embedded"`
	BodyParts   []CharacterBodyPart      `json:"BodyParts" gorm:"-"`
	UpdateTime  time.Time                `json:"UpdateTime"`
	Immortal    bool                     `json:"Immortal"`
}

type CharacterHealthAttribute struct {
	Current int `json:"Current"`
	Maximum int `json:"Maximum"`
}

type CharacterBodyPart struct {
	Name   string                   `json:"Name"`
	Health CharacterHealthAttribute `json:"Health" gorm:"embedded"`
}

type CharacterInventory struct {
	Items              []CharacterItem `json:"items" gorm:"-"`
	Equipment          string          `json:"equipment"`
	Stash              string          `json:"stash"`
	SortingTable       string          `json:"sortingTable"`
	QuestRaidItems     string          `json:"questRaidItems"`
	QuestStashItems    string          `json:"questStashItems"`
	FastPanel          datatypes.JSON  `json:"fastPanel"`
	HideoutAreaStashes datatypes.JSON  `json:"hideoutAreaStashes"`
	FavoriteItems      []string        `json:"favoriteItems"`
}

type CharacterItem struct {
	ID       string                `json:"_id"`
	Tpl      string                `json:"_tpl"`
	ParentID string                `json:"parentId"`
	SlotID   string                `json:"slotId"`
	Location CharacterItemLocation `json:"location"`
	Upd      CharacterItemUpd      `json:"upd"`
}

type CharacterItemLocation struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
}

type CharacterItemUpd struct {
	StackObjectsCount int                 `json:"stackObjectsCount"`
	Repairable        CharacterRepairable `json:"repairable"`
	SpawnedInSession  bool                `json:"spawnedInSession"`
}

type CharacterRepairable struct {
	Durability    int `json:"durability"`
	MaxDurability int `json:"maxDurability"`
}

type CharacterSkills struct {
	Common    []CharacterSkill `json:"Common"`
	Mastering []CharacterSkill `json:"Mastering"`
	Points    int              `json:"Points"`
}

type CharacterSkill struct {
	ID                        string  `json:"id"`
	Progress                  float64 `json:"progress"`
	PointsEarnedDuringSession int     `json:"pointsEarnedDuringSession"`
	LastAccess                int64   `json:"lastAccess"`
}

type CharacterStats struct {
	Eft CharacterEftStats `json:"eft"`
}

type CharacterEftStats struct {
	SessionCounters        CharacterSessionCounters `json:"SessionCounters"`
	OverallCounters        CharacterOverallCounters `json:"OverallCounters"`
	SessionExperienceMult  float64                  `json:"SessionExperienceMult"`
	ExperienceBonusMult    float64                  `json:"ExperienceBonusMult"`
	TotalSessionExperience int                      `json:"TotalSessionExperience"`
	LastSessionDate        time.Time                `json:"LastSessionDate"`
	Aggressor              string                   `json:"Aggressor"`
	DroppedItems           []string                 `json:"DroppedItems"`
	FoundInRaidItems       []string                 `json:"FoundInRaidItems"`
	Victims                []string                 `json:"Victims"`
	CarriedQuestItems      []string                 `json:"CarriedQuestItems"`
	DamageHistory          CharacterDamageHistory   `json:"DamageHistory"`
	LastPlayerState        string                   `json:"LastPlayerState"`
	TotalInGameTime        int                      `json:"TotalInGameTime"`
	SurvivorClass          string                   `json:"SurvivorClass"`
}

type CharacterSessionCounters struct {
	Items []CharacterCounterItem `json:"Items"`
}

type CharacterOverallCounters struct {
	Items []CharacterCounterItem `json:"Items"`
}

type CharacterCounterItem struct {
	Key   []string `json:"Key"`
	Value int      `json:"Value"`
}

type CharacterDamageHistory struct {
	LethalDamagePart string                               `json:"LethalDamagePart"`
	LethalDamage     string                               `json:"LethalDamage"`
	BodyParts        map[string][]CharacterBodyPartDamage `json:"BodyParts"`
}

type CharacterBodyPartDamage struct {
	Amount         int    `json:"Amount"`
	Type           string `json:"Type"`
	SourceID       string `json:"SourceId"`
	OverDamageFrom string `json:"OverDamageFrom"`
	Blunt          bool   `json:"Blunt"`
	ImpactsCount   int    `json:"ImpactsCount"`
}

type CharacterHideout struct {
	Production   datatypes.JSON         `json:"Production"`
	Areas        []CharacterHideoutArea `json:"Areas"`
	Improvements datatypes.JSON         `json:"Improvements"`
	Seed         int                    `json:"Seed"`
}

type CharacterHideoutArea struct {
	Type                  int              `json:"type"`
	Level                 int              `json:"level"`
	Active                bool             `json:"active"`
	PassiveBonusesEnabled bool             `json:"passiveBonusesEnabled"`
	CompleteTime          int              `json:"completeTime"`
	Constructing          bool             `json:"constructing"`
	Slots                 []datatypes.JSON `json:"slots"`
	LastRecipe            string           `json:"lastRecipe"`
}

type CharacterBonus struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	TemplateID string `json:"templateId"`
}

type CharacterNotes struct {
	Notes []datatypes.JSON `json:"Notes"`
}

type CharacterAchievements struct {
	Achievements map[string]int64 `json:"Achievements"`
}

type CharacterRagfairInfo struct {
	Rating          float64          `json:"rating"`
	IsRatingGrowing bool             `json:"isRatingGrowing"`
	Offers          []datatypes.JSON `json:"offers"`
}

type CharacterTradersInfo struct {
	TradersInfo map[string]CharacterTraderInfo `json:"TradersInfo"`
}

type CharacterTraderInfo struct {
	Unlocked bool    `json:"unlocked"`
	Disabled bool    `json:"disabled"`
	SalesSum int     `json:"salesSum"`
	Standing float64 `json:"standing"`
}

type CharacterUnlockedInfo struct {
	UnlockedProductionRecipe []datatypes.JSON `json:"unlockedProductionRecipe"`
}
