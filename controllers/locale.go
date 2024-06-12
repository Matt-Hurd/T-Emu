package controllers

import (
	"eft-private-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetLocaleEn(c *gin.Context) {
	helpers.ServeJSONFile("client.locale.en.json")(c)
}
