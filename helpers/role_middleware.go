package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateRoleMiddleware(role byte) func(*fiber.Ctx) error {
	// must be used after JWT middleware to include the claims
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		if claims["Role"] == float64(role) {
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"Error": "unauthorized",
			})
	}
}
