// websocket/server.go
package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"notifier-server/db"
	"notifier-server/models"

	"github.com/gorilla/websocket"
)

var Channels = make(map[string]map[*websocket.Conn]bool)
var ChannelsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request, channelID string) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	if channelID == "" {
		log.Println("ChannelID is required")
		return
	}

	ChannelsMutex.Lock()
	if Channels[channelID] == nil {
		Channels[channelID] = make(map[*websocket.Conn]bool)
	}
	Channels[channelID][ws] = true
	ChannelsMutex.Unlock()

	// Check for unsent notifications and send them
	var notifications []models.Notification
	db.DB.Where("channel_id = ?", channelID).Find(&notifications)
	for _, notification := range notifications {
		err := ws.WriteMessage(websocket.TextMessage, []byte(notification.Message))
		if err == nil {
			db.DB.Delete(&notification)
		}
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			ChannelsMutex.Lock()
			delete(Channels[channelID], ws)
			if len(Channels[channelID]) == 0 {
				delete(Channels, channelID)
			}
			ChannelsMutex.Unlock()
			break
		} else {
			fmt.Printf("Received message: %s\n", string(msg))
		}
	}
}
