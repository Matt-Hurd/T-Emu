package match

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentMatchGroupResponse struct {
	Squad []string `json:"squad"`
}

func GetMatchGroupCurrent(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", CurrentMatchGroupResponse{})
}
