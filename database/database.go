package database

import (
	"management-backend/config"

	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Unexported variable to implement singleton pattern
var db *gorm.DB

func Connect() {
	// Fetch config to get DB URI
	cfg := config.Get()
	if cfg.DBUri == "" {
		log.Fatal("DB URI not set in config")
	}

	// Open SQL connection
	pgConn, err := sql.Open("pgx", cfg.DBUri)
	if err != nil {
		log.Fatal(err)
	}

	// Open gorm connection
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: pgConn}))
	if err != nil {
		log.Fatal(err)
	}

	// Log success
	log.Println("Database connected!")
}

/*
Get will return the config set in Init
*/
func Get() *gorm.DB {
	return db
}
