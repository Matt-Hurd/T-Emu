package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetGlobals(c *gin.Context) {
	helpers.ServeJSONFile("client.globals.json")(c)
}
