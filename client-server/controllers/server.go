package controllers

import (
	"client-server/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetServerList(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", []map[string]interface{}{
		{
			"ip":   os.Getenv("GAME_SERVER_ADDRESS"),
			"port": os.Getenv("GAME_SERVER_PORT"),
		}, {
			"ip":   os.Getenv("GAME_SERVER_ADDRESS"),
			"port": os.Getenv("GAME_SERVER_PORT"),
		},
	})
}
