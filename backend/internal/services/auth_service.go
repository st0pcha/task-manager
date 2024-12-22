package services

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/internal/types"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
	"github.com/st0pcha/task-manager/backend/pkg/utils/checkers"
	"github.com/st0pcha/task-manager/backend/pkg/utils/password"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx, req *types.RegisterDTO) error {
	user := &dal.User{}
	if err := dal.FindUserByEmail(user, req.Email).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ErrorResponse(c, fiber.StatusConflict, "user with this email already exists")
	}

	log.Println(3)

	if err := checkers.ValidateEmail(req.Email); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := checkers.ValidatePasswordsMatch(req.Password, req.PasswordRepeat); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := checkers.ValidatePassword(req.Password); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	hashedPassword := password.Hash(req.Password)
	createdUser := &dal.User{
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := dal.CreateUser(createdUser).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "failed to create user")
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "user registered", &types.UserResponse{
		Email: createdUser.Email,
		ID:    createdUser.ID.String(),
	})
}

func LoginUser(c *fiber.Ctx, req *types.AuthDTO) error {
	user := &dal.User{}
	if err := dal.FindUserByEmail(user, req.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "user with this email not exists")
		}
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "failed to find user")
	}
	if !password.Verify(user.Password, req.Password) {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid credentials")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "user logged in", &types.UserResponse{
		Email: user.Email,
		ID:    user.ID.String(),
	})
}
