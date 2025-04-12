package models

type Role struct {
	ID   uint   `json:"role_id" gorm:"primaryKey;column:role_id"`
	Name string `json:"role_name" gorm:"column:role_name"`
}

// TableName указывает имя таблицы для модели
func (Role) TableName() string {
	return "role"
}
