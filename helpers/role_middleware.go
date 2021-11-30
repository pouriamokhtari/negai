package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewRoleMiddleware(role string) func(*fiber.Ctx) error {
	// must be used after JWT middleware to include the claims
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		log.Println(claims["Role"].(string))
		if claims["Role"].(string) == role {
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"Error": "unauthorized",
			})
	}
}
