package user

import (
	"log"

	"management-backend/config"
	"management-backend/database"
	"management-backend/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

/*
AddRoutes adds routes for all user endpoints to given router
*/
func AddRoutes(router fiber.Router) {
	// Fetch DB and config, and check if they've been set
	db := database.Get()
	if db == nil {
		log.Fatalln("Connect to DB before adding routes")
	}
	cfg := config.Get()
	if cfg == nil {
		log.Fatalln("Read config variables before adding routes")
	}

	// Create new user service
	h := Handler{
		DB:     db,
		Config: cfg,
	}

	// Set routes for all POST requests
	router.Post("/", h.Create)

	// Set routes for all GET requests
	router.Get("/:id", jwt.JwtHandler(cfg, h.Read))

	// Set routes for all PATCH requests
	router.Patch("/", jwt.JwtHandler(cfg, h.Update))

	// Set routes for all DELETE requests
	router.Delete("/:id", jwt.JwtHandler(cfg, h.Delete))
}
