package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RaidConfigurationRequest struct {
	KeyId                  string `json:"keyId"`
	Location               string `json:"location"`
	TimeVariant            string `json:"timeVariant"`
	MetabolismDisabled     bool   `json:"metabolismDisabled"`
	TimeAndWeatherSettings struct {
		IsRandomTime    bool   `json:"isRandomTime"`
		IsRandomWeather bool   `json:"isRandomWeather"`
		CloudinessType  string `json:"cloudinessType"`
		RainType        string `json:"rainType"`
		WindType        string `json:"windType"`
		FogType         string `json:"fogType"`
		TimeFlowType    string `json:"timeFlowType"`
		HourOfDay       int    `json:"hourOfDay"`
	} `json:"timeAndWeatherSettings"`
	BotSettings struct {
		IsScavWars bool   `json:"isScavWars"`
		BotAmount  string `json:"botAmount"`
	} `json:"botSettings"`
	WavesSettings struct {
		BotAmount         string `json:"botAmount"`
		BotDifficulty     string `json:"botDifficulty"`
		IsBosses          bool   `json:"isBosses"`
		IsTaggedAndCursed bool   `json:"isTaggedAndCursed"`
	} `json:"wavesSettings"`
	Side                string `json:"side"`
	RaidMode            string `json:"raidMode"`
	PlayersSpawnPlace   string `json:"playersSpawnPlace"`
	CanShowGroupPreview bool   `json:"CanShowGroupPreview"`
	MaxGroupCount       int    `json:"MaxGroupCount"`
}

func GetRaidConfiguration(c *gin.Context) {
	var req RaidConfigurationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}
	helpers.JSONResponse(c, http.StatusOK, "", nil)
}
