package helpers

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

var JWTMiddleware func(*fiber.Ctx) error

func NewJWTMiddleware(config jwt.Config) {
	JWTMiddleware = jwt.New(config)
}
