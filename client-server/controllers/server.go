package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerList(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", []map[string]interface{}{
		{
			"ip":   "127.0.0.1",
			"port": 9090,
		},
	})
}
