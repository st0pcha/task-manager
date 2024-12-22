package response

import "github.com/gofiber/fiber/v3"

func ErrorResponse(ctx fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"data":    nil,
	})
}

func SuccessResponse(ctx fiber.Ctx, code int, message string, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}
