package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetLocaleEn(c *gin.Context) {
	helpers.ServeJSONFile("client.locale.en.json")(c)
}
