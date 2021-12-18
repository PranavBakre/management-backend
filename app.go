package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"management-backend/api"
	"management-backend/config"
	"management-backend/database"
	mgmtError "management-backend/utils/error"
	"management-backend/utils/googleapis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":8000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func init() {
	// Configure logging
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stderr)

	// Fetch config vars
	config.Init()

	//Init googleApiClient
	cfg := config.Get()
	googleapis.Init(
		cfg.ClientId,
		cfg.ClientSecret,
		cfg.RedirectUri,
		"authorization_code")

	// Connected with database
	database.Connect()
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			e, ok := err.(mgmtError.Error)
			if ok {
				code = e.Code
				return c.Status(code).JSON(e)
			}
			return c.Status(code).JSON(map[string]string{
				"code":    strconv.Itoa(code),
				"message": "An unexpected error has occurred",
			})
		}})

	// Middleware
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())

	// Add API routes to /api endpoint
	api.AddRoutes(app.Group("/api"))

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
