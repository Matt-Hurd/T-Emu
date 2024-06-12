package hideout

import (
	"client-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetHideoutAreas(c *gin.Context) {
	helpers.ServeJSONFile("client.hideout.areas.json")(c)
}

func GetHideoutQteList(c *gin.Context) {
	helpers.ServeJSONFile("client.hideout.qte.list.json")(c)
}

func GetHideoutSettings(c *gin.Context) {
	helpers.ServeJSONFile("client.hideout.settings.json")(c)
}

func GetHideoutProductionRecipes(c *gin.Context) {
	helpers.ServeJSONFile("client.hideout.production.recipes.json")(c)
}

func GetHideoutProductionScavcaseRecipes(c *gin.Context) {
	helpers.ServeJSONFile("client.hideout.production.scavcase.recipes.json")(c)
}
