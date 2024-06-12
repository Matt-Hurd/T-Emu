package game

import (
	"client-server/config"
	"client-server/helpers"
	"client-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileList(c *gin.Context) {
	id := config.GetSession(c).AccountID
	var user models.Account

	// Query the database with preloading related data
	if err := config.DB.Preload("Profiles").
		Preload("Profiles.Info").
		Preload("Profiles.Info.Settings").
		Preload("Profiles.Customization").
		Preload("Profiles.Health").
		Preload("Profiles.Health.BodyParts").
		Preload("Profiles.Inventory.Items").
		Preload("Profiles.Skills.Skills").
		Preload("Profiles.Stats.Eft.SessionCounters.Items").
		Preload("Profiles.Stats.Eft.OverallCounters.Items").
		Preload("Profiles.Stats.Eft.DamageHistory.BodyParts").
		Preload("Profiles.Hideout.Areas").
		Preload("Profiles.Bonuses").
		Preload("Profiles.Notes").
		Preload("Profiles.Quests").
		Preload("Profiles.Achievements").
		Preload("Profiles.RagfairInfo").
		Preload("Profiles.WishList").
		Preload("Profiles.TradersInfo").
		Preload("Profiles.TradersInfo.Traders").
		Preload("Profiles.UnlockedInfo").
		Preload("Profiles.Status").
		First(&user, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", user.Profiles)
}

type ProfileSelectRequest struct {
	Uid string `json:"uid"`
}

type ProfileSelectResponseData struct {
	Status string `json:"status"`
}

func ProfileSelect(c *gin.Context) {
	var req ProfileSelectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	var profile models.Profile
	if err := config.DB.Where("id = ? AND account_id = ?", req.Uid, config.GetSession(c).AccountID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.JSONResponse(c, http.StatusBadRequest, "Invalid profile", nil)
			return
		}
		helpers.JSONResponse(c, http.StatusBadRequest, "Unknown Error", nil)
		return
	}

	if !config.UpdateSessionProfileID(config.GetSessionID(c), profile.ID) {
		helpers.JSONResponse(c, http.StatusBadRequest, "Failed to update session", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", ProfileSelectResponseData{Status: "ok"})
}
