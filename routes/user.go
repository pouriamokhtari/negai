package routes

import (
	"negai/handlers"
	"negai/helpers"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

func BindUser(r fiber.Router) {
	// JWT middleware
	r.Use(helpers.JWTMiddleware)
	// Role middleware
	r.Use(helpers.NewRoleMiddleware(
		helpers.RoleMiddlewareConfig{
			Role:    models.Admin,
			Handler: handlers.Unauthorized,
		}))
	r.Get("/", handlers.GetAllUsers)
	r.Get("/:id", handlers.GetUser)
	r.Post("/", handlers.CreateUser)
	r.Patch("/:id", handlers.UpdateUser)
	r.Delete("/:id", handlers.DeleteUser)
}
