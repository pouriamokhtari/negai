package routes

import (
	"negai/handlers"
	"negai/helpers"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

func BindUser(r fiber.Router) {
	// JWT Middleware
	r.Use(helpers.JWTMiddleware)
	r.Use(helpers.NewRoleMiddleware(models.Admin))
	r.Get("/", handlers.GetAllUsers)
	r.Get("/:id", handlers.GetUser)
	r.Post("/", handlers.CreateUser)
	r.Patch("/:id", handlers.UpdateUser)
	r.Delete("/:id", handlers.DeleteUser)
}
