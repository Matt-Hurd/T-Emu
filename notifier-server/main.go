// main.go
package main

import (
	"log"
	"net/http"

	"notifier-server/db"
	"notifier-server/handlers"
	"notifier-server/websocket"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db.Init()

	router := gin.Default()

	router.POST("/create-channel", handlers.CreateChannel)
	router.DELETE("/delete-channel/:channel_id", handlers.DeleteChannel)
	router.POST("/send-notification", handlers.SendNotification)
	router.GET("/push/notifier/get/:channelid", func(c *gin.Context) {
		websocket.HandleConnections(c.Writer, c.Request)
	})

	log.Fatal(http.ListenAndServe(":8090", router))
}
