package services

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/internal/types"
	"github.com/st0pcha/task-manager/backend/pkg/utils"
	"gorm.io/gorm"
)

func GetSelfUserWithTasks(c *fiber.Ctx) error {
	user, err := utils.GetSelf(c)
	if err != nil {
		return err
	}
	tasks := []dal.Task{}
	if err := dal.FindTasksByUserID(&tasks, user.ID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "failed to find tasks")
		}
	}

	tasksRes := make([]types.TaskResponse, len(tasks))
	for i, task := range tasks {
		tasksRes[i] = types.TaskResponse{
			ID:      task.ID.String(),
			Content: task.Content,
			IsDone:  task.IsDone,
		}
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "user found", &types.UserWithTasksResponse{
		User:  types.UserResponse{Email: user.Email, ID: user.ID.String()},
		Tasks: tasksRes,
	})
}
