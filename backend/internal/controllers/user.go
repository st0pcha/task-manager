package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/services"
)

func GetSelfUserWithTasks(c *fiber.Ctx) error {
	return services.GetSelfUserWithTasks(c)
}
