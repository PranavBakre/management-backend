package api

import (
	"github.com/PranavBakre/management-backend/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
BindAPI mounts all the individual APIs to the main router
*/
func BindAPI(router fiber.Router, db *gorm.DB, cfg config.Config) {
	// Mount individual APIs to prefixes
}
