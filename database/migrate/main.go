package main

import (
	"log"
	"os"

	"management-backend/config"
	"management-backend/database"
	"management-backend/models"
)

func init() {
	// Configure logging
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stderr)

	// Fetch config vars
	config.Init()

	// Connected with database
	database.Connect()
}

func main() {
	db := database.Get()

	// Add UUID extension to postgres
	result := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Migrate all models
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
