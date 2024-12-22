package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, data interface{}) *fiber.Error {
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid validation error: %v", err))
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, vErr := range validationErrors {
				errors = append(errors, fmt.Sprintf("Field '%s': %s", vErr.Field(), vErr.Tag()))
			}
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Validation failed: %s", errors))
		}

		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unexpected validation error: %v", err))
	}

	return nil
}
