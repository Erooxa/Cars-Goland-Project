package models

type User struct {
	ID       uint   `json:"user_id" gorm:"primaryKey;column:user_id"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id" gorm:"column:role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}
