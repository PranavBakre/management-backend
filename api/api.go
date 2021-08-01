package api

import (
	"github.com/PranavBakre/management-backend/api/user"
	"github.com/PranavBakre/management-backend/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
AddRoutes adds routes for all the individual APIs to the given router
*/
func AddRoutes(router fiber.Router, db *gorm.DB, cfg config.Config) {
	user.AddRoutes(router.Group("/user"), db, cfg)
}
