package user

import (
	"github.com/PranavBakre/management-backend/config"
	"github.com/PranavBakre/management-backend/database"

	"log"

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
	svc := Service{
		DB:     db,
		Config: cfg,
	}

	// Set routes for all POST requests
	router.Post("/", svc.Create)

	// Set routes for all GET requests
	router.Get("/", svc.ReadAll)
	router.Get("/:id", svc.Read)

	// Set routes for all PATCH requests
	router.Patch("/", svc.Update)

	// Set routes for all DELETE requests
	router.Delete("/:id", svc.Delete)
}
