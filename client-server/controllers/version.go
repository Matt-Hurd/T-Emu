package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCheckVersion(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", map[string]interface{}{
		"isvalid":       true,
		"latestversion": "1.14.8.5.30150",
	})
}
