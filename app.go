package main

import (
	"negai/database"
	"negai/handlers"
	"negai/middlewares"
	"negai/models"
	"negai/routes"

	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwt "github.com/gofiber/jwt/v3"
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
	// migrate models
	models.AutoMigrateModels()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork:      *prod, // go run app.go -prod
		ErrorHandler: handlers.InternalServerError,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create JWT middleware (used later with route groups)
	middlewares.NewJWTMiddleware(jwt.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: handlers.InvalidJWT,
	})

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
