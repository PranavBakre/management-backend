package utils

import (
	"management-backend/config"
	"github.com/golang-jwt/jwt"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/m4rw3r/uuid"
)

/*
JwtHandler returns given handler wrapped in JWT middleware
*/
func JwtHandler(cfg *config.Config, fn fiber.Handler) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ContextKey:     "jwt",
		SigningKey:     []byte(cfg.JwtSecret),
		SuccessHandler: fn,
	})
}

/*
CreateToken creates a JWT token for given user ID
*/
func CreateToken(cfg *config.Config, id uuid.UUID) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	// Generate encoded token
	tkn, err := token.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tkn, nil
}

/*
GetCurrentUserID will return the user ID set in the JWT token present in passed context
*/
func GetCurrentUserID(ctx *fiber.Ctx) (uuid.UUID, error) {
	token := ctx.Locals("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return uuid.FromString(claims["id"].(string))
}
