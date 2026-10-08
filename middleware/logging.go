package middleware

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

type LoggingMiddleware struct{}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Handle(c telebot.Context) error {
	start := time.Now()
	user := c.Sender()

	log.Printf(
		"[REQUEST] User: %s (%d) | Chat: %d | Message: %s",
		user.Username,
		user.ID,
		c.Chat().ID,
		c.Text(),
	)

	err := c.Next()

	duration := time.Since(start)
	log.Printf(
		"[RESPONSE] Duration: %v | Error: %v",
		duration,
		err,
	)

	return err
}
