package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/st0pcha/task-manager/backend/pkg/response"
)

func GetHelloWorld(ctx fiber.Ctx) error {
	return response.SuccessResponse(ctx, fiber.StatusOK, "Hello world!", nil)
}
