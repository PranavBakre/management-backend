package database

import (
	"github.com/PranavBakre/management-backend/config"

	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Connect() *gorm.DB {
	if db == nil {
		cfg := config.GetConfig()
		if cfg.DBUri == "" {
			log.Fatal("DB URI not set in config")
		}

		pgConn, err := sql.Open("pgx", cfg.DBUri)
		if err != nil {
			log.Fatal(err)
		}

		db, err = gorm.Open(postgres.New(postgres.Config{Conn: pgConn}))
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug()

		log.Println("Database connected!")
	}

	return db
}
