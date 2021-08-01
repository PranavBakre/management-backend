package main

import (
	"github.com/PranavBakre/management-backend/api"
	"github.com/PranavBakre/management-backend/config"
	"github.com/PranavBakre/management-backend/database"

	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Configure logging
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stderr)

	// Parse command-line flags
	flag.Parse()

	// Fetch config vars
	cfg := config.GetConfig()

	// Connected with database
	db := database.Connect(cfg)

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Add API routes to /api endpoint
	api.AddRoutes(app.Group("/api"), db, cfg)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
