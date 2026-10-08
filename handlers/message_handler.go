package handlers

import (
	"telegram-bot/database"
	"telegram-bot/models"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type MessageHandler struct {
	DB *gorm.DB
}

func NewMessageHandler(db *gorm.DB) *MessageHandler {
	return &MessageHandler{DB: db}
}

func (h *MessageHandler) HandleTextMessage(c telebot.Context) error {
	user := c.Sender()
	text := c.Text()

	// Guardar mensaje en la base de datos
	message := models.Message{
		TelegramID:  c.Message().ID,
		UserID:      uint64(user.ID),
		ChatID:      c.Chat().ID,
		Text:        text,
		MessageType: "text",
		IsFromBot:   false,
	}

	if err := h.DB.Create(&message).Error; err != nil {
		return c.Send("❌ Error al guardar mensaje")
	}

	// Procesar el mensaje y generar respuesta
	response := h.processMessage(text)

	// Guardar respuesta del bot
	botMessage := models.Message{
		TelegramID:  0, // Se asignará después de enviar
		UserID:      uint64(user.ID),
		ChatID:      c.Chat().ID,
		Text:        response,
		MessageType: "text",
		IsFromBot:   true,
	}

	sentMsg, err := c.Send(response)
	if err != nil {
		return err
	}

	botMessage.TelegramID = sentMsg.ID
	h.DB.Create(&botMessage)

	return nil
}

func (h *MessageHandler) processMessage(text string) string {
	// Lógica simple de procesamiento de mensajes
	// Aquí puedes integrar IA, NLP, etc.
	
	switch text {
	case "hola", "Hola", "HOLA":
		return "¡Hola! ¿En qué puedo ayudarte?"
	case "adiós", "Adiós", "ADIOS":
		return "¡Hasta luego! 👋"
	default:
		return "He recibido tu mensaje: \"" + text + "\""
	}
}

func (h *MessageHandler) HandlePhoto(c telebot.Context) error {
	user := c.Sender()
	photo := c.Message().Photo

	// Guardar información de la foto
	message := models.Message{
		TelegramID:  c.Message().ID,
		UserID:      uint64(user.ID),
		ChatID:      c.Chat().ID,
		Text:        "[Foto: " + photo.FileID + "]",
		MessageType: "photo",
		IsFromBot:   false,
	}

	h.DB.Create(&message)

	return c.Send("📷 He recibido tu foto")
}

func (h *MessageHandler) HandleDocument(c telebot.Context) error {
	user := c.Sender()
	doc := c.Message().Document

	message := models.Message{
		TelegramID:  c.Message().ID,
		UserID:      uint64(user.ID),
		ChatID:      c.Chat().ID,
		Text:        "[Documento: " + doc.FileName + "]",
		MessageType: "document",
		IsFromBot:   false,
	}

	h.DB.Create(&message)

	return c.Send("📄 He recibido tu documento")
}

func (h *MessageHandler) HandleCallback(c telebot.Context) error {
	callback := c.Callback()
	
	// Procesar callback
	return c.Respond(&telebot.CallbackResponse{
		Text: "✅ Acción completada",
	})
}
