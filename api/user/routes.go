package user

import (
	"github.com/PranavBakre/management-backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
AddRoutes adds routes for all user endpoints to given router
*/
func AddRoutes(router fiber.Router, db *gorm.DB, cfg config.Config) {
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
