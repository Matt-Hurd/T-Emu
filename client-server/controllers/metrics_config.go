package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetMetricsConfig(c *gin.Context) {
	helpers.ServeJSONFile("client.getMetricsConfig.json")(c)
}
