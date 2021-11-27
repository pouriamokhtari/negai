package helpers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var AuthMiddleware func(*fiber.Ctx) error

func CreateAuthMiddleware() {
	AuthMiddleware = jwtware.New(jwtware.Config{
		SigningKey: os.Getenv("JWT_SECRET"),
	})
}
