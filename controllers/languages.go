package controllers

import (
	"eft-private-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var LANGUAGES = map[string]string{
	"ch":    "Chinese",
	"cz":    "Czech",
	"en":    "English",
	"fr":    "French",
	"ge":    "German",
	"hu":    "Hungarian",
	"it":    "Italian",
	"jp":    "Japanese",
	"kr":    "Korean",
	"pl":    "Polish",
	"po":    "Portugal",
	"sk":    "Slovak",
	"es":    "Spanish",
	"es-mx": "Spanish Mexico",
	"tu":    "Turkish",
	"ru":    "Русский",
	"ro":    "Romanian",
}

func GetLanguages(c *gin.Context) {
	helpers.JSONResponse(c, http.StatusOK, "", LANGUAGES)
}
