package models

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	ChannelID     string         `gorm:"uniqueIndex"`
	Notifications []Notification `gorm:"foreignKey:ChannelID;references:ChannelID"`
}
