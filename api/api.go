package api

import (
	"github.com/PranavBakre/management-backend/api/user"

	"github.com/gofiber/fiber/v2"
)

/*
AddRoutes adds routes for all the individual APIs to the given router
*/
func AddRoutes(router fiber.Router) {
	user.AddRoutes(router.Group("/user"))
}
