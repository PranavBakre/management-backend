package user

import (
	"github.com/PranavBakre/management-backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
API creates and returns a fiber app for the user service
*/
func API(db *gorm.DB, cfg config.Config) *fiber.App {
	// Create new app and user service
	app := fiber.New(fiber.Config{})
	svc := Service{
		DB:     db,
		Config: cfg,
	}

	// Set routes for all POST requests
	app.Post("/", svc.Create)

	// Set routes for all GET requests
	app.Get("/", svc.ReadAll)
	app.Get("/:id", svc.Read)

	// Set routes for all PATCH requests
	app.Patch("/", svc.Update)

	// Set routes for all DELETE requests
	app.Delete("/:id", svc.Delete)

	return app
}
