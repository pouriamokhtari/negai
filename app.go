package main

import (
	"negai/database"
	"negai/handlers"
	"negai/helpers"
	"negai/routes"
	"os"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect(os.Getenv("DATABASE_CONNECTION"))
	database.AutoMigrateModels()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create JWT middleware (used later with route groups)
	helpers.CreateAuthMiddleware()

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind routes
	routes.BindUser(v1.Group("/user"))
	routes.BindAuth(v1.Group("/auth"))

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
