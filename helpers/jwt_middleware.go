package helpers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

var JWTMiddleware func(*fiber.Ctx) error

func NewJWTMiddleware() {
	JWTMiddleware = jwt.New(jwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}
