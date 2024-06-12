package friend

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFriendList(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", map[string]interface{}{
		"Friends":      []string{},
		"Ignore":       []string{},
		"InIgnoreList": []string{},
	})
}
