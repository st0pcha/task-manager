package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	return nil
}

func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(c, body); err != nil {
		return err
	}
	return Validate(c, body)
}
func ParseAllowedOrigins(origins string) string {
	if origins == "*" {
		return "*"
	}
	return strings.ReplaceAll(origins, ",", ", ")
}
