package controllers

import (
	"eft-private-server/config"
	"eft-private-server/helpers"
	"eft-private-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileCharacters(c *gin.Context) {
	id := "1"
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
