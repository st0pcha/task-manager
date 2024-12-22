package dal

import (
	"github.com/st0pcha/task-manager/backend/internal/db"
	"gorm.io/gorm"
)

type User struct {
	Base

	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Tasks    []Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tasks"`
}

func CreateUser(user *User) *gorm.DB {
	return db.DB.Create(user)
}

func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.DB.Model(&User{}).Take(dest, conds...)
}

func FindUserByID(dest interface{}, id string) *gorm.DB {
	return FindUser(dest, "id = ?", &id)
}

func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", &email)
}
