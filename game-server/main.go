package main

import (
	"log"

	"game-server/server"
)

func main() {
	srv := server.NewServer()
	err := srv.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
