package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"data":    nil,
	})
}

func SuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}
