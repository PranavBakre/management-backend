package database

import (
	"github.com/PranavBakre/management-backend/config"

	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	if cfg.DBUri == "" {
		log.Fatal("DB URI not set in config")
	}

	// Open SQL connection
	pgConn, err := sql.Open("pgx", cfg.DBUri)
	if err != nil {
		log.Fatal(err)
	}

	// Open gorm connection
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: pgConn}))
	if err != nil {
		log.Fatal(err)
	}

	// Log success
	log.Println("Database connected!")

	return db
}
