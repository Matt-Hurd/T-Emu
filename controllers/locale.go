package controllers

import (
	"eft-private-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetEnLocale(c *gin.Context) {
	helpers.ServeJSONFile("client.menu.locale.en.json")(c)
}
