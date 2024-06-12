package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetAchievementStatistic(c *gin.Context) {
	helpers.ServeJSONFile("client.achievement.statistic.json")(c)
}

func GetAchievementList(c *gin.Context) {
	helpers.ServeJSONFile("client.achievement.list.json")(c)
}
