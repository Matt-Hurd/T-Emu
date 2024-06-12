package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetQuestList(c *gin.Context) {
	helpers.ServeJSONFile("client.quest.list.json")(c)
}
