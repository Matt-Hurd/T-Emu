package menu

import (
	"eft-private-server/helpers"

	"github.com/gin-gonic/gin"
)

func GetMenuLocaleEn(c *gin.Context) {
	helpers.ServeJSONFile("client.menu.locale.en.json")(c)
}
