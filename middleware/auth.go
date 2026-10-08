package middleware

import (
	"telegram-bot/database"
	"telegram-bot/models"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	DB *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{DB: db}
}

func (m *AuthMiddleware) Handle(c telebot.Context) error {
	user := c.Sender()

	var dbUser models.User
	result := m.DB.Where("telegram_id = ?", user.ID).First(&dbUser)

	if result.Error != nil {
		return c.Send("❌ Debes iniciar el bot primero con /start")
	}

	if !dbUser.IsActive {
		return c.Send("❌ Tu cuenta está desactivada")
	}

	return c.Next()
}
