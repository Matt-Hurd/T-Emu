package models

import "encoding/json"

type ProfileInfo struct {
	BaseModel
	ProfileID               string `json:"-"`
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
	Settings                ProfileInfoSettings
	NicknameChangeDate      int64 `gorm:"default:0"`
	NeedWipeOptions         []ProfileInfoNeedWipeOption
	LastCompletedWipe       *int  `gorm:"default:null"`
	LastCompletedEvent      *int  `gorm:"default:null"`
	BannedState             bool  `gorm:"default:false"`
	BannedUntil             int64 `gorm:"default:0"`
	IsStreamerModeAvailable bool  `gorm:"default:false"`
	SquadInviteRestriction  bool  `gorm:"default:false"`
	HasCoopExtension        bool  `gorm:"default:false"`
	Bans                    []ProfileBan
}

type ProfileInfoNeedWipeOption struct {
	ProfileInfoID string `gorm:"primaryKey" json:"-"`
	WipeOptionID  string `gorm:"primaryKey"`
}

type ProfileBan struct {
	ProfileInfoID string `gorm:"primaryKey" json:"-"`
	BanID         string `gorm:"primaryKey"`
}

type ProfileInfoSettings struct {
	ProfileInfoID   string  `gorm:"primaryKey" json:"-"`
	Role            string  `gorm:"default:'assault'"`
	BotDifficulty   string  `gorm:"default:'easy'"`
	Experience      int     `gorm:"default:-1"`
	StandingForKill float64 `gorm:"default:0"`
	AggressorBonus  float64 `gorm:"default:0"`
}

func (ci ProfileInfo) MarshalJSON() ([]byte, error) {
	type Alias ProfileInfo
	if ci.NeedWipeOptions == nil {
		ci.NeedWipeOptions = []ProfileInfoNeedWipeOption{}
	}
	if ci.Bans == nil {
		ci.Bans = []ProfileBan{}
	}
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(ci),
	})
}
