package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetHandbookBuildsMyList(c *gin.Context) {
	helpers.ServeJSONFile("client.builds.list.json")(c)
}
