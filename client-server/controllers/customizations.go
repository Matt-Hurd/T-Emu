package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetCustomization(c *gin.Context) {
	helpers.ServeJSONFile("client.customization.json")(c)
}
