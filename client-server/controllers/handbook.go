package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetHandbookTemplates(c *gin.Context) {
	helpers.ServeJSONFile("client.handbook.templates.json")(c)
}
