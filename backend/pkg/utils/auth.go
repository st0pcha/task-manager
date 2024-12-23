package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/pkg/utils/jwt"
	"gorm.io/gorm"
)

func GetSelf(c *fiber.Ctx) (*dal.User, error) {
	refreshToken := c.Cookies(jwt.RefreshToken)
	if refreshToken == "" {
		return nil, ErrorResponse(c, fiber.StatusUnauthorized, "refresh token is required")
	}

	claims, err := jwt.ValidateJWT(refreshToken)
	if err != nil {
		return nil, ErrorResponse(c, fiber.StatusUnauthorized, "refresh token is invalid")
	}

	userID := claims["sub"].(string)
	user := &dal.User{}
	if err := dal.FindUserByID(user, userID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrorResponse(c, fiber.StatusInternalServerError, "can't find user")
	}

	return user, nil
}
