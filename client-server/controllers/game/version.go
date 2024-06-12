package game

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionValidateVersion struct {
	Major    string `json:"major"`
	Minor    string `json:"minor"`
	Game     string `json:"game"`
	Backend  string `json:"backend"`
	Taxonomy string `json:"taxonomy"`
}

type VersionValidateRequest struct {
	Version VersionValidateVersion `json:"version"`
	Develop bool                   `json:"develop"`
}

func VersionValidate(c *gin.Context) {
	var req VersionValidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	if req.Version.Major != "0.14.9.0.30473" || req.Version.Minor != "live" || req.Version.Game != "live" || req.Version.Backend != "6" || req.Version.Taxonomy != "341" {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid game version", nil)
		return
	}

	// Send the JSON response
	helpers.JSONResponse(c, http.StatusOK, "", nil)
}
