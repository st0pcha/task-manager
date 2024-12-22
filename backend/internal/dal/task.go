package dal

import (
	"github.com/google/uuid"
)

type Task struct {
	Base

	Content string    `json:"content"`
	IsDone  bool      `gorm:"default:false" json:"is_done"`
	UserID  uuid.UUID `gorm:"not null" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID" json:"user"`
}
