package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ChannelID string `gorm:"index"`
	Message   string
}
