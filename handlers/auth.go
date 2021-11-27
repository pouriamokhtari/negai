package handlers

import (
	"negai/database"
	"negai/helpers"
	"negai/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type RegisterParams struct {
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

func Register(c *fiber.Ctx) error {
	params := &RegisterParams{}

	if err := c.BodyParser(params); err != nil {
		return BadRequest(c)
	}

	passwordDigest, err := helpers.HashPassword(params.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Email:          params.Email,
		PasswordDigest: passwordDigest,
	}
	result := database.Connection.Create(&user)

	if result.Error != nil {
		return BadRequest(c)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	params := &LoginParams{}

	if err := c.BodyParser(params); err != nil {
		return BadRequest(c)
	}

	user := &models.User{}

	result := database.Connection.Where("email = ?", params.Email).First(&user)
	if result.Error != nil {
		return Unauthorized(c)
	}

	if !helpers.CheckPasssword(params.Password, user.PasswordDigest) {
		return Unauthorized(c)
	}

	claims := jwt.MapClaims{
		"Email":   user.Email,
		"Expires": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": encodedToken,
	})
}
