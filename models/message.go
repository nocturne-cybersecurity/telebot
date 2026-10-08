package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	TelegramID  int64          `gorm:"not null;index" json:"telegram_id"`
	UserID      uint64         `gorm:"not null;index" json:"user_id"`
	ChatID      int64          `gorm:"not null;index" json:"chat_id"`
	Text        string         `gorm:"type:text" json:"text"`
	MessageType string         `gorm:"size:50" json:"message_type"`
	IsFromBot   bool           `gorm:"default:false" json:"is_from_bot"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Message) TableName() string {
	return "messages"
}
