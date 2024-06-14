package match

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMatchAvailable(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", true)
}
