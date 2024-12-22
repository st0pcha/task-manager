package dal

type User struct {
	Base

	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Tasks    []Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tasks"`
}
