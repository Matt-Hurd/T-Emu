// handlers/notifications.go
package handlers

import (
	"net/http"

	"notifier-server/db"
	"notifier-server/models"
	"notifier-server/websocket"

	"github.com/gin-gonic/gin"
	gorillawebsocket "github.com/gorilla/websocket"
)

type NotificationRequest struct {
	ChannelID string `json:"channel_id" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

func SendNotification(c *gin.Context) {
	var req NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	websocket.ChannelsMutex.Lock()
	subscribers, exists := websocket.Channels[req.ChannelID]
	websocket.ChannelsMutex.Unlock()

	if exists {
		success := false
		for subscriber := range subscribers {
			err := subscriber.WriteMessage(gorillawebsocket.TextMessage, []byte(req.Message))
			if err == nil {
				success = true
			}
		}
		if success {
			c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
			return
		}
	}

	notification := models.Notification{
		ChannelID: req.ChannelID,
		Message:   req.Message,
	}
	db.DB.Create(&notification)

	c.JSON(http.StatusOK, gin.H{"status": "notification stored"})
}
