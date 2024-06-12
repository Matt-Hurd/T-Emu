package friend

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFriendRequestListOutbox(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", []string{})
}

func GetFriendRequestListInbox(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", []string{})
}
