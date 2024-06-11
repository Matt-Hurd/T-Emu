package main

import (
	"eft-private-server/config"
	"eft-private-server/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	config.ConnectDatabase()

	// Migrate the database
	config.MigrateDatabase()

	// Set up and run the router
	r := routes.SetupRouter()
	r.Run(":8080") // Run on port 8080
}
