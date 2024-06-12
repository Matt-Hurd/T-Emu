package controllers

import (
	"database/sql"
	"eft-private-server/config"
	"eft-private-server/helpers"
	"eft-private-server/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

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

func createDefaultTraderInfo() []models.CharacterTraderInfo {
	defaultTraders := []string{"5ac3b934156ae10c4430e83c",
		"58330581ace78e27b8b10cee",
		"656f0f98d80a697f855d34b1",
		"5c0647fdd443bc2504c2d371",
		"5a7c2eca46aef81a7ca2145d",
		"5935c25fb3acc3127c3d8cd9",
		"579dc571d53a0658a154fbec",
		"54cb50c76803fa8b248b4571",
		"638f541a29ffd1183d187f57",
		"54cb57776803fa99248b456e",
	}
	var defaultTraderInfo []models.CharacterTraderInfo
	for _, trader := range defaultTraders {
		defaultTraderInfo = append(defaultTraderInfo, models.CharacterTraderInfo{
			TraderID: trader,
			Unlocked: false,
			Disabled: false,
			SalesSum: 0,
			Standing: 0,
		})
	}
	return defaultTraderInfo
}

func BuildNewPlayerCharacter(accountID int) *models.Character {
	equipment := &models.CharacterItem{
		ID:  helpers.GenerateUUID(),
		Tpl: "55d7217a4bdc2d86028b456d",
	}
	questRaidItems := &models.CharacterItem{
		ID:  helpers.GenerateUUID(),
		Tpl: "5963866286f7747bf429b572",
	}
	questStashItems := &models.CharacterItem{
		ID:  helpers.GenerateUUID(),
		Tpl: "5963866b86f7747bfa1c4462",
	}
	stash := &models.CharacterItem{
		ID:  helpers.GenerateUUID(),
		Tpl: "5811ce772459770e9e5f9532",
	}
	sortingTable := &models.CharacterItem{
		ID:  helpers.GenerateUUID(),
		Tpl: "602543c13fee350cd564d032",
	}
	newCharacter := &models.Character{
		ID:        helpers.GenerateUUID(),
		AccountID: accountID,
		Info: models.CharacterInfo{
			Nickname:         "NewPlayer",
			LowerNickname:    "newplayer",
			Side:             "Usec",
			Voice:            "Usec_4",
			RegistrationDate: time.Now().Unix(),
			GameVersion:      "edge_of_darkness_with_pve",
			Settings: models.CharacterInfoSettings{
				Role:            "assault",
				BotDifficulty:   "easy",
				Experience:      -1,
				StandingForKill: 0,
				AggressorBonus:  0,
			},
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
			Items:              []models.CharacterItem{*equipment, *questRaidItems, *stash, *questStashItems, *sortingTable},
			FastPanel:          datatypes.JSON([]byte("{}")),
			HideoutAreaStashes: datatypes.JSON([]byte("{}")),
			FavoriteItems:      datatypes.JSON([]byte("[]")),
			Equipment:          (*models.CharacterItemID)(&equipment.ID),
			QuestRaidItems:     (*models.CharacterItemID)(&questRaidItems.ID),
			Stash:              (*models.CharacterItemID)(&stash.ID),
			QuestStashItems:    (*models.CharacterItemID)(&questStashItems.ID),
			SortingTable:       (*models.CharacterItemID)(&sortingTable.ID),
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
				LastSessionDate:        time.Now().Unix(),
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
		Bonuses: []models.CharacterBonus{
			{ID: helpers.GenerateUUID(), Type: "StashSize", TemplateID: "566abbc34bdc2d92178b4576"},
			{ID: helpers.GenerateUUID(), Type: "StashSize", TemplateID: "5811ce572459770cba1a34ea"},
			{ID: helpers.GenerateUUID(), Type: "StashSize", TemplateID: "5811ce662459770f6f490f32"},
			{ID: helpers.GenerateUUID(), Type: "StashSize", TemplateID: "5811ce772459770e9e5f9532"},
		},
		Notes:        models.CharacterNotes{Notes: datatypes.JSON([]byte("[]"))},
		Quests:       []models.CharacterQuest{},
		Achievements: []models.CharacterAchievement{},
		RagfairInfo: models.CharacterRagfairInfo{
			Rating:          0.2,
			IsRatingGrowing: true,
			Offers:          datatypes.JSON([]byte("[]")),
		},
		WishList: []models.CharacterWishListItem{},
		TradersInfo: models.CharacterTradersInfo{
			Traders: createDefaultTraderInfo(),
		},
		UnlockedInfo: models.CharacterUnlockedInfo{UnlockedProductionRecipe: datatypes.JSON([]byte("[]"))},
	}

	return newCharacter
}

func BuildNewScavCharacter(accountID int) *models.Character {
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
			Nickname:           "Миха Телега",
			LowerNickname:      "Серый Снайпер",
			Side:               "Savage",
			Voice:              "Scav_2",
			GameVersion:        "edge_of_darkness_with_pve",
			LockedMoveCommands: true,
			RegistrationDate:   time.Now().Unix(),
			Settings: models.CharacterInfoSettings{
				Role:            "assault",
				BotDifficulty:   "normal",
				Experience:      -1,
				StandingForKill: -0.1,
				AggressorBonus:  0.03,
			},
			NeedWipeOptions: []models.CharacterInfoNeedWipeOption{},
			Bans:            []models.CharacterBan{},
		},
		Customization: models.CharacterCustomization{
			Head:  "5f68c4c217d579077152a252",
			Body:  "5fd1eb3fbe3b7107d66cb645",
			Feet:  "5f5e410c6bdad616ad46d60b",
			Hands: "5fd78fe9e3bfcf6cab4c9f54",
		},
		Health: models.CharacterHealth{
			Hydration:   models.CharacterHealthAttribute{Current: 100, Maximum: 100},
			Energy:      models.CharacterHealthAttribute{Current: 100, Maximum: 100},
			Temperature: models.CharacterHealthAttribute{Current: 36.6, Maximum: 40},
			BodyParts: []models.CharacterBodyPart{
				{Name: "Head", Health: models.CharacterHealthAttribute{Current: 35, Maximum: 35}},
				{Name: "Chest", Health: models.CharacterHealthAttribute{Current: 85, Maximum: 85}},
				{Name: "Stomach", Health: models.CharacterHealthAttribute{Current: 70, Maximum: 70}},
				{Name: "LeftArm", Health: models.CharacterHealthAttribute{Current: 60, Maximum: 60}},
				{Name: "RightArm", Health: models.CharacterHealthAttribute{Current: 60, Maximum: 60}},
				{Name: "LeftLeg", Health: models.CharacterHealthAttribute{Current: 65, Maximum: 65}},
				{Name: "RightLeg", Health: models.CharacterHealthAttribute{Current: 65, Maximum: 65}},
			},
			UpdateTime: time.Now().Unix(),
		},
		Inventory: models.CharacterInventory{
			Items:              []models.CharacterItem{*equipment, *questRaidItems},
			FastPanel:          datatypes.JSON([]byte("{}")),
			HideoutAreaStashes: datatypes.JSON([]byte("{}")),
			FavoriteItems:      datatypes.JSON([]byte("[]")),
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
				LastSessionDate:        time.Now().Unix(),
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

	return newCharacter
}

func CreateCharacter(c *gin.Context) {
	accountID, atoi_err := strconv.Atoi(c.Param("accountID"))

	if atoi_err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, atoi_err.Error(), nil)
		return
	}

	newCharacter := BuildNewPlayerCharacter(accountID)
	scavCharacter := BuildNewScavCharacter(accountID)

	newCharacter.SavageID = &scavCharacter.ID

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newCharacter).Error; err != nil {
			return err
		}
		if err := tx.Create(scavCharacter).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "", newCharacter)
}
