package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
	"github.com/st0pcha/task-manager/backend/pkg/utils/jwt"
)

func IsAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "missing token")
	}
	token := parts[1]

	user, err := jwt.ParseJWT(token)
	if err != nil {
		if err.Error() == "token is expired" {
			utils.ClearCookie(c, jwt.RefreshToken)
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "failed to parse token")
	}
	c.Locals("userID", user.ID)
	return c.Next()
}
