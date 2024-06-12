package models

import "encoding/json"

type CharacterInfo struct {
	BaseModel
	CharacterID             string `json:"-"`
	Nickname                string
	LowerNickname           string
	Side                    string
	Voice                   string
	Level                   int   `gorm:"default:1"`
	Experience              int   `gorm:"default:0"`
	RegistrationDate        int64 `gorm:"default:0"`
	GameVersion             string
	AccountType             int   `gorm:"default:0"`
	MemberCategory          int   `gorm:"default:2"`
	LockedMoveCommands      bool  `gorm:"default:false" json:"lockedMoveCommands"`
	SavageLockTime          int64 `gorm:"default:0"`
	LastTimePlayedAsSavage  int64 `gorm:"default:0"`
	Settings                CharacterInfoSettings
	NicknameChangeDate      int64 `gorm:"default:0"`
	NeedWipeOptions         []CharacterInfoNeedWipeOption
	LastCompletedWipe       *int  `gorm:"default:null"`
	LastCompletedEvent      *int  `gorm:"default:null"`
	BannedState             bool  `gorm:"default:false"`
	BannedUntil             int64 `gorm:"default:0"`
	IsStreamerModeAvailable bool  `gorm:"default:false"`
	SquadInviteRestriction  bool  `gorm:"default:false"`
	HasCoopExtension        bool  `gorm:"default:false"`
	Bans                    []CharacterBan
}

type CharacterInfoNeedWipeOption struct {
	CharacterInfoID string `gorm:"primaryKey" json:"-"`
	WipeOptionID    string `gorm:"primaryKey"`
}

type CharacterBan struct {
	CharacterInfoID string `gorm:"primaryKey" json:"-"`
	BanID           string `gorm:"primaryKey"`
}

type CharacterInfoSettings struct {
	CharacterInfoID string  `gorm:"primaryKey" json:"-"`
	Role            string  `gorm:"default:'assault'"`
	BotDifficulty   string  `gorm:"default:'easy'"`
	Experience      int     `gorm:"default:-1"`
	StandingForKill float64 `gorm:"default:0"`
	AggressorBonus  float64 `gorm:"default:0"`
}

func (ci CharacterInfo) MarshalJSON() ([]byte, error) {
	type Alias CharacterInfo
	if ci.NeedWipeOptions == nil {
		ci.NeedWipeOptions = []CharacterInfoNeedWipeOption{}
	}
	if ci.Bans == nil {
		ci.Bans = []CharacterBan{}
	}
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(ci),
	})
}
