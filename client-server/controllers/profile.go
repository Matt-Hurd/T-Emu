package controllers

import (
	"client-server/config"
	"client-server/helpers"
	"client-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileStatusRequest struct {
	MaxPveCountExceeded bool                   `json:"maxPveCountExceeded"`
	Profiles            []models.ProfileStatus `json:"profiles"`
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

	var profiles []models.ProfileStatus
	for _, profile := range user.Profiles {
		profiles = append(profiles, profile.Status)
	}

	response := ProfileStatusRequest{
		MaxPveCountExceeded: false,
		Profiles:            profiles,
	}

	helpers.JSONResponse(c, http.StatusOK, "", response)

}
