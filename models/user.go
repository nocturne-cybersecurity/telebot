package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	TelegramID   int64          `gorm:"uniqueIndex;not null" json:"telegram_id"`
	Username     string         `gorm:"size:255" json:"username"`
	FirstName    string         `gorm:"size:255" json:"first_name"`
	LastName     string         `gorm:"size:255" json:"last_name"`
	LanguageCode string         `gorm:"size:10" json:"language_code"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	IsAdmin      bool           `gorm:"default:false" json:"is_admin"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type UserState struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null;index" json:"user_id"`
	State     string    `gorm:"size:100" json:"state"`
	Data      string    `gorm:"type:text" json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (UserState) TableName() string {
	return "user_states"
}
