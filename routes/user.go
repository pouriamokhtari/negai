package routes

import (
	"negai/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func BindUser(r fiber.Router) {
	// JWT Middleware
	r.Use(jwtware.New(jwtware.Config{
		SigningKey: os.Getenv("JWT_SECRET"),
	}))
	r.Get("/", handlers.GetAllUsers)
	r.Get("/:id", handlers.GetUser)
	r.Post("/", handlers.CreateUser)
	r.Patch("/:id", handlers.UpdateUser)
	r.Delete("/:id", handlers.DeleteUser)
}
