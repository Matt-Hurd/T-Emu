// main.go
package main

import (
	"log"
	"net/http"

	"notifier-server/controllers"
	"notifier-server/db"
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

	router.POST("/create-channel", controllers.CreateChannel)
	router.DELETE("/delete-channel/:channel_id", controllers.DeleteChannel)
	router.POST("/send-notification", controllers.SendNotification)
	router.GET("/push/notifier/getwebsocket/:channelid", func(c *gin.Context) {
		channelID := c.Param("channelid")
		websocket.HandleConnections(c.Writer, c.Request, channelID)
	})

	log.Fatal(http.ListenAndServe(":8090", router))
}
