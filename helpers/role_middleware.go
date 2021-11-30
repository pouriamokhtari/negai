package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type RoleMiddlewareConfig struct {
	Handler func(*fiber.Ctx) error
	Role    string
}

func NewRoleMiddleware(config RoleMiddlewareConfig) func(*fiber.Ctx) error {
	// must be used after JWT middleware to include the claims
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		if claims["Role"].(string) == config.Role {
			return c.Next()
		}
		return config.Handler(c)
	}
}
