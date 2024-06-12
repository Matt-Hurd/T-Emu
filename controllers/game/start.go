package game

import (
	"eft-private-server/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GameStartResponseData struct {
	UtcTime float64 `json:"utc_time"`
}

func GameStart(c *gin.Context) {
	respData := GameStartResponseData{
		UtcTime: float64(time.Now().UnixNano()) / 1e9,
	}

	// Send the JSON response
	helpers.JSONResponse(c, http.StatusOK, "", respData)
}
