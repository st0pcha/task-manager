package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.StatusOK, "Hello world!", nil)
}
