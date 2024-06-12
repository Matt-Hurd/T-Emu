package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRepeatableQuestsActivityPeriods(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", []string{})
}
