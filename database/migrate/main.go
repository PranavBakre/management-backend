package main

import (
	"log"
	"os"

	"management-backend/config"
	"management-backend/database"
	"management-backend/models"

	"gorm.io/gorm/clause"
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
	cfg := config.Get()
	// Add UUID extension to postgres
	result := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	// Migrate all models
	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Right{},
	)

	var godRight = models.Right{
		Right: "GOD",
	}
	result = db.FirstOrCreate(&godRight)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	var godRole = models.Role{
		Role:   "GOD",
		Rights: []models.Right{godRight},
	}

	result = db.FirstOrCreate(&godRole)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	var god = models.User{
		GoogleID: &cfg.SuperUserGoogleId,
		Name:     cfg.SuperUserName,
		Email:    cfg.SuperUserEmail,
		Roles:    []models.Role{godRole},
	}
	result = db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "google_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"google_id", "name", "email"}),
		}).
		Create(&god)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if err != nil {
		log.Fatal(err)
	}
}
