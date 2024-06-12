package controllers

import (
	"database/sql"
	"eft-private-server/config"
	"eft-private-server/helpers"
	"eft-private-server/models"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func GetCharacters(c *gin.Context) {
	id := c.Param("id")
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
		Preload("Characters.UnlockedInfo").
		First(&user, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", user)
}

func createDefaultSkillsCommon() []models.CharacterSkill {
	commonSkills := []string{"BotReload",
		"BotSound",
		"Endurance",
		"Strength",
		"Vitality",
		"Health",
		"StressResistance",
		"Metabolism",
		"Immunity",
		"Perception",
		"Intellect",
		"Attention",
		"Charisma",
		"Pistol",
		"Revolver",
		"SMG",
		"Assault",
		"Shotgun",
		"Sniper",
		"LMG",
		"HMG",
		"Launcher",
		"AttachedLauncher",
		"Throwing",
		"Melee",
		"DMR",
		"AimDrills",
		"TroubleShooting",
		"Surgery",
		"CovertMovement",
		"Search",
		"MagDrills",
		"FieldMedicine",
		"FirstAid",
		"LightVests",
		"HeavyVests",
		"NightOps",
		"SilentOps",
		"WeaponTreatment",
		"Auctions",
		"Cleanoperations",
		"Shadowconnections",
		"Taskperformance",
		"Crafting",
		"HideoutManagement",
	}
	var defaultCommonSkills []models.CharacterSkill
	zero := 0.0
	var minInt int32 = math.MinInt32
	for _, skill := range commonSkills {
		defaultCommonSkills = append(defaultCommonSkills, models.CharacterSkill{
			ID:                        skill,
			Type:                      "Common",
			Progress:                  0,
			PointsEarnedDuringSession: &zero,
			LastAccess:                &minInt,
		})
	}
	return defaultCommonSkills
}

func CreateCharacter(c *gin.Context) {
	accountID := c.Param("accountID")
	equipment := &models.CharacterItem{
		ID:               helpers.GenerateUUID(),
		Tpl:              "55d7217a4bdc2d86028b456d",
		SpawnedInSession: sql.NullBool{Bool: true, Valid: true},
	}
	questRaidItems := &models.CharacterItem{
		ID:               helpers.GenerateUUID(),
		Tpl:              "5963866286f7747bf429b572",
		SpawnedInSession: sql.NullBool{Bool: true, Valid: true},
	}
	newCharacter := &models.Character{
		ID:        helpers.GenerateUUID(),
		AccountID: accountID,
		Info: models.CharacterInfo{
			Nickname:        "NewPlayer",
			LowerNickname:   "newplayer",
			Side:            "USEC",
			Voice:           "Default",
			GameVersion:     "0.12.9",
			Settings:        models.CharacterInfoSettings{},
			NeedWipeOptions: []models.CharacterInfoNeedWipeOption{},
			Bans:            []models.CharacterBan{},
		},
		Customization: models.CharacterCustomization{
			Head:  "5cde96047d6c8b20b577f016",
			Body:  "5cde95d97d6c8b647a3769b0",
			Feet:  "5cde95ef7d6c8b04713c4f2d",
			Hands: "5cde95fa7d6c8b04737c2d13",
		},
		Health: models.CharacterHealth{
			Hydration:   models.CharacterHealthAttribute{Current: 100, Maximum: 100},
			Energy:      models.CharacterHealthAttribute{Current: 100, Maximum: 100},
			Temperature: models.CharacterHealthAttribute{Current: 36, Maximum: 40},
			BodyParts: []models.CharacterBodyPart{
				{Name: "Head", Health: models.CharacterHealthAttribute{Current: 35, Maximum: 35}},
				{Name: "Chest", Health: models.CharacterHealthAttribute{Current: 85, Maximum: 85}},
				{Name: "Stomach", Health: models.CharacterHealthAttribute{Current: 70, Maximum: 70}},
				{Name: "LeftArm", Health: models.CharacterHealthAttribute{Current: 60, Maximum: 60}},
				{Name: "RightArm", Health: models.CharacterHealthAttribute{Current: 60, Maximum: 60}},
				{Name: "LeftLeg", Health: models.CharacterHealthAttribute{Current: 65, Maximum: 65}},
				{Name: "RightLeg", Health: models.CharacterHealthAttribute{Current: 65, Maximum: 65}},
			},
		},
		Inventory: models.CharacterInventory{
			Items:              []models.CharacterItem{*equipment, *questRaidItems},
			FastPanel:          datatypes.JSON([]byte("{}")),
			HideoutAreaStashes: datatypes.JSON([]byte("{}")),
			FavoriteItems:      []models.CharacterInventoryFavoriteItem{},
			Equipment:          (*models.CharacterItemID)(&equipment.ID),
			QuestRaidItems:     (*models.CharacterItemID)(&questRaidItems.ID),
		},
		Skills: models.CharacterSkillsGroup{
			Skills: createDefaultSkillsCommon(),
			Points: 0,
		},
		Stats: models.CharacterStats{
			Eft: models.CharacterEftStats{
				SessionCounters:        models.CharacterCounter{Items: []models.CharacterCounterItem{}},
				OverallCounters:        models.CharacterCounter{Items: []models.CharacterCounterItem{}},
				SessionExperienceMult:  0,
				ExperienceBonusMult:    1.0,
				TotalSessionExperience: 0,
				LastSessionDate:        time.Now(),
				Aggressor:              nil,
				DroppedItems:           []string{},
				FoundInRaidItems:       []string{},
				Victims:                []string{},
				CarriedQuestItems:      []string{},
				DamageHistory:          models.CharacterDamageHistory{BodyParts: []models.CharacterBodyPartDamage{}},
				LastPlayerState:        nil,
				TotalInGameTime:        0,
				SurvivorClass:          "Survivor",
			},
		},
		TaskConditionCounters: datatypes.JSON([]byte("{}")),
		InsuredItems:          datatypes.JSON([]byte("[]")),
		Hideout: models.CharacterHideout{
			Production:   datatypes.JSON([]byte("{}")),
			Areas:        []models.CharacterHideoutArea{},
			Improvements: datatypes.JSON([]byte("{}")),
			Seed:         0,
		},
		Bonuses:      []models.CharacterBonus{},
		Notes:        models.CharacterNotes{Notes: datatypes.JSON([]byte("[]"))},
		Quests:       []models.CharacterQuest{},
		Achievements: []models.CharacterAchievement{},
		RagfairInfo: models.CharacterRagfairInfo{
			Rating:          0.2,
			IsRatingGrowing: true,
			Offers:          datatypes.JSON([]byte("[]")),
		},
		WishList:     []models.CharacterWishListItem{},
		TradersInfo:  models.CharacterTradersInfo{},
		UnlockedInfo: models.CharacterUnlockedInfo{UnlockedProductionRecipe: datatypes.JSON([]byte("[]"))},
	}

	if err := config.DB.Create(newCharacter).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helpers.JSONResponse(c, http.StatusOK, "", newCharacter)
}
