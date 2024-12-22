package dal

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
