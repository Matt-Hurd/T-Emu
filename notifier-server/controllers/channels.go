package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"notifier-server/db"
	"notifier-server/models"

	"github.com/gin-gonic/gin"
)

func generateChannelID() (string, error) {
	bytes := make([]byte, 32) // 32 bytes will result in 64 characters after encoding
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateChannel(c *gin.Context) {
	channelID, err := generateChannelID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate channel ID"})
		return
	}

	channel := models.Channel{ChannelID: channelID}
	if err := db.DB.Create(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create channel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"channel_id": channelID})
}
func DeleteChannel(c *gin.Context) {
	channelID := c.Param("channel_id")
	if channelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ChannelID is required"})
		return
	}

	if err := db.DB.Where("channel_id = ?", channelID).Delete(&models.Channel{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete channel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Channel deleted"})
}
