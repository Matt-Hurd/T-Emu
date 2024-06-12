package game

import (
	"eft-private-server/config"
	"eft-private-server/helpers"
	"eft-private-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileList(c *gin.Context) {
	id := config.GetSession(c).AccountID
	var user models.Account

	// Query the database with preloading related data
	if err := config.DB.Preload("Characters").
		Preload("Characters.Info").
		Preload("Characters.Info.Settings").
		Preload("Characters.Customization").
		Preload("Characters.Health").
		Preload("Characters.Health.BodyParts").
		Preload("Characters.Inventory.Items").
		Preload("Characters.Skills.Skills").
		Preload("Characters.Stats.Eft.SessionCounters.Items").
		Preload("Characters.Stats.Eft.OverallCounters.Items").
		Preload("Characters.Stats.Eft.DamageHistory.BodyParts").
		Preload("Characters.Hideout.Areas").
		Preload("Characters.Bonuses").
		Preload("Characters.Notes").
		Preload("Characters.Quests").
		Preload("Characters.Achievements").
		Preload("Characters.RagfairInfo").
		Preload("Characters.WishList").
		Preload("Characters.TradersInfo").
		Preload("Characters.TradersInfo.Traders").
		Preload("Characters.UnlockedInfo").
		First(&user, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", user.Characters)
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

	var character models.Character
	if err := config.DB.Where("id = ? AND account_id = ?", req.Uid, config.GetSession(c).AccountID).First(&character).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.JSONResponse(c, http.StatusBadRequest, "Invalid character", nil)
			return
		}
		helpers.JSONResponse(c, http.StatusBadRequest, "Unknown Error", nil)
		return
	}

	if !config.UpdateSessionProfileID(config.GetSessionID(c), character.ID) {
		helpers.JSONResponse(c, http.StatusBadRequest, "Failed to update session", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", ProfileSelectResponseData{Status: "ok"})
}
