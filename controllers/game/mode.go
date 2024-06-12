package game

import (
	"eft-private-server/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type GetGameModeRequest struct {
	SessionMode interface{} `json:"sessionMode"`
}

type GetGameModeResponseData struct {
	GameMode   string `json:"gameMode"`
	BackendUrl string `json:"backendUrl"`
}

func GetGameMode(c *gin.Context) {
	var req GetGameModeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	respData := GetGameModeResponseData{
		GameMode:   "pve",
		BackendUrl: "https://" + os.Getenv("SERVER_URL"),
	}

	// Send the JSON response
	helpers.JSONResponse(c, http.StatusOK, "", respData)
}
