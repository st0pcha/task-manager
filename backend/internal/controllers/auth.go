package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/internal/services"
	"github.com/st0pcha/task-manager/backend/internal/types"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
	"github.com/st0pcha/task-manager/backend/pkg/utils/checkers"
)

func Register(c *fiber.Ctx) error {
	var req types.RegisterDTO

	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	var userCheck dal.User
	if err := dal.FindUserByEmail(&userCheck, req.Email).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "user with this email already exists")
	}

	if err := checkers.ValidateEmail(req.Email); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := checkers.ValidatePasswordsMatch(req.Password, req.PasswordRepeat); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := checkers.ValidatePassword(req.Password); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := services.RegisterUser(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "user registered", &types.UserResponse{
		Email: user.Email,
		ID:    user.ID.String(),
	})
}
