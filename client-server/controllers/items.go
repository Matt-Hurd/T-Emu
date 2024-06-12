package controllers

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	helpers.ServeJSONFile("client.items.json")(c)
}

func GetItemPrices(c *gin.Context) {
	helpers.ServeJSONFile("client.items.prices.json")(c)
}
