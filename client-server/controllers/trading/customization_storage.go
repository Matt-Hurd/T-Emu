package trading

import (
	"client-server/config"
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomizationStorageResponseData struct {
	Id     string   `json:"_id"`
	Suites []string `json:"suites"`
}

func GetCustomizationStorage(c *gin.Context) {
	profileId := config.GetSession(c).ProfileID
	respData := CustomizationStorageResponseData{
		Id:     profileId,
		Suites: []string{"5cde9ec17d6c8b04723cf479", "5cde9e957d6c8b0474535da7"},
	}

	helpers.JSONResponse(c, http.StatusOK, "", respData)
}
