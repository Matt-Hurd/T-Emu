package controllers

import (
	"bytes"
	"client-server/config"
	"client-server/helpers"
	"client-server/models"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateChannelResponse struct {
	ChannelID string `json:"channel_id"`
}

type CreateNotifierChannelResponse struct {
	Server         string `json:"server"`
	ChannelID      string `json:"channel_id"`
	URL            string `json:"url"`
	NotifierServer string `json:"notifierServer"`
	WS             string `json:"ws"`
}

func CreateNotifierChannel(c *gin.Context) {
	notifierServerURL := os.Getenv("NOTIFIER_SERVER_URL")
	url := "http://" + notifierServerURL + "/create-channel"

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	var createChannelResp CreateChannelResponse
	if err := json.NewDecoder(resp.Body).Decode(&createChannelResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	var profile models.Profile
	id := config.GetSession(c).ProfileID
	if err := config.DB.First(&profile, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "Profile not found", nil)
		return
	}
	profile.NotificationChannel = createChannelResp.ChannelID
	if err := config.DB.Updates(&profile).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to update profile notification channel", nil)
		return
	}

	respData := CreateNotifierChannelResponse{
		Server:         notifierServerURL,
		ChannelID:      createChannelResp.ChannelID,
		URL:            "",
		NotifierServer: "http://" + notifierServerURL + "/push/notifier/get/" + createChannelResp.ChannelID,
		WS:             "ws://" + notifierServerURL + "/push/notifier/getwebsocket/" + createChannelResp.ChannelID,
	}

	helpers.JSONResponse(c, http.StatusOK, "", respData)
}
