package handlers

import (
	"context"
	"fmt"
	"telegram-bot/database"
	"telegram-bot/models"

	"gopkg.in/telebot.v3"
)

type CommandHandler struct {
	DB *gorm.DB
}

func NewCommandHandler(db *gorm.DB) *CommandHandler {
	return &CommandHandler{DB: db}
}

func (h *CommandHandler) HandleStart(c telebot.Context) error {
	user := c.Sender()

	// Crear o actualizar usuario en la base de datos
	var dbUser models.User
	result := h.DB.Where("telegram_id = ?", user.ID).First(&dbUser)

	if result.Error != nil {
		// Crear nuevo usuario
		dbUser = models.User{
			TelegramID:   user.ID,
			Username:     user.Username,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			LanguageCode: user.LanguageCode,
			IsActive:     true,
		}
		if err := h.DB.Create(&dbUser).Error; err != nil {
			return c.Send("❌ Error al registrar usuario")
		}
	} else {
		// Actualizar usuario existente
		dbUser.Username = user.Username
		dbUser.FirstName = user.FirstName
		dbUser.LastName = user.LastName
		dbUser.IsActive = true
		h.DB.Save(&dbUser)
	}

	welcomeMsg := fmt.Sprintf(
		"👋 ¡Hola %s!\n\n"+
			"Bienvenido al bot. Los comandos disponibles son:\n"+
			"/start - Iniciar el bot\n"+
			"/help - Mostrar ayuda\n"+
			"/about - Información del bot\n"+
			"/settings - Configuración",
		user.FirstName,
	)

	return c.Send(welcomeMsg)
}

func (h *CommandHandler) HandleHelp(c telebot.Context) error {
	helpMsg := "📚 *Comandos disponibles:*\n\n"+
		"/start - Iniciar el bot\n"+
		"/help - Mostrar esta ayuda\n"+
		"/about - Información sobre el bot\n"+
		"/settings - Configuración del usuario\n"+
		"/stats - Estadísticas de uso\n\n"+
		"Para cualquier duda, contacta al administrador."

	return c.Send(helpMsg, telebot.ModeMarkdown)
}

func (h *CommandHandler) HandleAbout(c telebot.Context) error {
	aboutMsg := "ℹ️ *Información del Bot*\n\n"+
		"Versión: 1.0.0\n"+
		"Desarrollado con Go y Telebot\n"+
		"Base de datos: PostgreSQL\n"+
		"Cache: Redis"

	return c.Send(aboutMsg, telebot.ModeMarkdown)
}

func (h *CommandHandler) HandleSettings(c telebot.Context) error {
	settingsMsg := "⚙️ *Configuración*\n\n"+
		"Opciones disponibles:\n"+
		"1. Cambiar idioma\n"+
		"2. Notificaciones\n"+
		"3. Privacidad\n\n"+
		"Próximamente..."

	return c.Send(settingsMsg, telebot.ModeMarkdown)
}

func (h *CommandHandler) HandleStats(c telebot.Context) error {
	var userCount int64
	var messageCount int64

	h.DB.Model(&models.User{}).Count(&userCount)
	h.DB.Model(&models.Message{}).Count(&messageCount)

	statsMsg := fmt.Sprintf(
		"📊 *Estadísticas*\n\n"+
			"👥 Usuarios: %d\n"+
			"💬 Mensajes: %d",
		userCount, messageCount,
	)

	return c.Send(statsMsg, telebot.ModeMarkdown)
}
