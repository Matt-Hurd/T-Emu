package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCheckVersion(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", map[string]interface{}{
		"isvalid":       true,
		"latestversion": "0.14.9.1.30626",
	})
}
