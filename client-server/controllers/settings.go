package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context) {
	helpers.ServeJSONFile("client.settings.json")(c)
}
