package routes

import (
	"negai/handlers"

	"github.com/gofiber/fiber/v2"
)

func BindAuth(r fiber.Router) {
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
}
