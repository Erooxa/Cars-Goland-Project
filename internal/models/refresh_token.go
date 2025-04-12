package models

import "time"

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token" gorm:"size:512;uniqueIndex"`
	UserID    uint      `json:"user_id" gorm:"column:user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
