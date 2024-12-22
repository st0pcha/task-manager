package services

import (
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/internal/types"
	"github.com/st0pcha/task-manager/backend/pkg/utils/password"
)

func RegisterUser(req *types.RegisterDTO) (user *dal.User, err error) {
	hashedPassword := password.Hash(req.Password)
	req.Password = hashedPassword
	createdUser := &dal.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := dal.CreateUser(createdUser).Error; err != nil {
		return nil, err
	}
	return createdUser, nil
}
