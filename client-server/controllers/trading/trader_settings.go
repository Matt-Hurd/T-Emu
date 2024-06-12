package trading

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetTraderSettings(c *gin.Context) {
	helpers.ServeJSONFile("client.trading.api.traderSettings.json")(c)
}
