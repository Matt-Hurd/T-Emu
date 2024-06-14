package controllers

import (
	"client-server/config"
	"client-server/helpers"
	"client-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LimitedProfileStatus struct {
	ProfileID    string  `gorm:"primaryKey" json:"profileid"`
	ProfileToken *string `json:"profileToken"`
	Status       string  `json:"status" gorm:"default:Free"`
	ServerId     string  `json:"sid"`
	IP           string  `json:"ip"`
	Port         int     `json:"port"`
}

type ProfileStatusRequest struct {
	MaxPveCountExceeded bool          `json:"maxPveCountExceeded"`
	Profiles            []interface{} `json:"profiles"`
}

func GetProfileStatus(c *gin.Context) {
	id := config.GetSession(c).AccountID
	var user models.Account

	if err := config.DB.Preload("Profiles").
		Preload("Profiles.Status").
		First(&user, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	var profiles []interface{}
	for _, profile := range user.Profiles {
		if profile.Status.Status == "Free" {
			profiles = append(profiles, LimitedProfileStatus{
				ProfileID:    profile.ID,
				ProfileToken: profile.Status.ProfileToken,
				Status:       profile.Status.Status,
				ServerId:     profile.Status.ServerId,
				IP:           profile.Status.IP,
				Port:         profile.Status.Port,
			})
		} else {
			profiles = append(profiles, profile.Status)
		}
	}

	response := ProfileStatusRequest{
		MaxPveCountExceeded: false,
		Profiles:            profiles,
	}

	helpers.JSONResponse(c, http.StatusOK, "", response)
}
