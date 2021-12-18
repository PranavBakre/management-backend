package auth

import (
	"log"
	"management-backend/config"
	"management-backend/database"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router) {
	db := database.Get()
	if db == nil {
		log.Fatalln("Connect to DB before adding routes")
	}
	cfg := config.Get()
	if cfg == nil {
		log.Fatalln("Read config variables before adding routes")
	}

	h := Handler{
		DB:     db,
		Config: cfg,
	}

	router.Post("/login", h.Login)
}
