package controllers

import (
	"eft-private-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {
	helpers.ServeJSONFile("client.locations.json")(c)
}
