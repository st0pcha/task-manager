package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/services"
	"github.com/st0pcha/task-manager/backend/internal/types"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
)

func Register(c *fiber.Ctx) error {
	var req types.RegisterDTO
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}
	return services.RegisterUser(c, &req)
}

func Login(c *fiber.Ctx) error {
	var req types.AuthDTO
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}
	return services.LoginUser(c, &req)
}
