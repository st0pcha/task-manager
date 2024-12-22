package dal

import (
	"github.com/google/uuid"
	"github.com/st0pcha/task-manager/backend/internal/db"
	"gorm.io/gorm"
)

type Task struct {
	Base

	Content string    `json:"content"`
	IsDone  bool      `gorm:"default:false" json:"is_done"`
	UserID  uuid.UUID `gorm:"not null" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID" json:"user"`
}

func CreateTask(task *Task) *gorm.DB {
	return db.DB.Create(task)
}

func FindTask(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.DB.Model(&Task{}).Take(dest, conds...)
}

func FindTaskByID(dest interface{}, id string) *gorm.DB {
	return FindTask(dest, "id = ?", &id)
}

func FindTasksByUserID(dest interface{}, userID uuid.UUID) *gorm.DB {
	return FindTask(dest, "user_id = ?", &userID)
}

func DeleteTaskByID(id string) *gorm.DB {
	return db.DB.Where("id = ?", id).Delete(&Task{})
}
