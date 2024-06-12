package game

import (
	"client-server/helpers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var LANGUAGES = map[string]string{
	"ch":    "Chinese",
	"cz":    "Czech",
	"en":    "English",
	"fr":    "French",
	"ge":    "German",
	"hu":    "Hungarian",
	"it":    "Italian",
	"jp":    "Japanese",
	"kr":    "Korean",
	"pl":    "Polish",
	"po":    "Portugal",
	"sk":    "Slovak",
	"es":    "Spanish",
	"es-mx": "Spanish Mexico",
	"tu":    "Turkish",
	"ru":    "Русский",
	"ro":    "Romanian",
}

type GameConfigRequest struct {
	AccountID         int               `json:"aid"`
	Lang              string            `json:"lang"`
	Languages         map[string]string `json:"languages"`
	NdaFree           bool              `json:"ndaFree"`
	Taxonomy          int               `json:"taxonomy"`
	Backend           map[string]string `json:"backend"`
	UseProtobuf       bool              `json:"useProtobuf"`
	UtcTime           float64           `json:"utc_time"`
	TotalInGame       int               `json:"totalInGame"`
	ReportAvailable   bool              `json:"reportAvailable"`
	TwitchEventMember bool              `json:"twitchEventMember"`
	SessionMode       string            `json:"sessionMode"`
}

func GetGameConfig(c *gin.Context) {
	serverUrl := os.Getenv("CLIENT_SERVER_URL")
	respData := GameConfigRequest{
		AccountID: 1,
		Lang:      "en",
		Languages: LANGUAGES,
		NdaFree:   false,
		Taxonomy:  6,
		Backend: map[string]string{
			"Lobby":     "ws://" + serverUrl + "/sws",
			"Trading":   "https://" + serverUrl,
			"Messaging": "https://" + serverUrl,
			"Main":      "https://" + serverUrl,
			"RagFair":   "https://" + serverUrl,
		},
		UseProtobuf:       false,
		UtcTime:           float64(time.Now().UnixNano()) / 1e9,
		TotalInGame:       2242554,
		ReportAvailable:   true,
		TwitchEventMember: false,
		SessionMode:       "pve",
	}

	helpers.JSONResponse(c, http.StatusOK, "", respData)
}
