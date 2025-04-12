package models

type Car struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Year        int     `json:"year"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
