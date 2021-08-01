package main

import (
	"github.com/PranavBakre/management-backend/config"
	"github.com/PranavBakre/management-backend/database"
	"github.com/PranavBakre/management-backend/models"

	"log"
)

func main() {
	cfg := config.GetConfig()
	db := database.Connect(cfg)

	// Add UUID extension to postgres
	result := db.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Scan(nil)
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
