package middleware

import (
	"fmt"
	"telegram-bot/database"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/telebot.v3"
)

type RateLimitMiddleware struct {
	Redis *redis.Client
}

func NewRateLimitMiddleware(redis *redis.Client) *RateLimitMiddleware {
	return &RateLimitMiddleware{Redis: redis}
}

func (m *RateLimitMiddleware) Handle(c telebot.Context) error {
	user := c.Sender()
	key := fmt.Sprintf("rate_limit:%d", user.ID)

	ctx := context.Background()

	// Verificar si el usuario ha excedido el límite
	count, err := m.Redis.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return c.Next()
	}

	if count >= 10 { // Límite de 10 mensajes por minuto
		return c.Send("⚠️ Has excedido el límite de mensajes. Espera un momento.")
	}

	// Incrementar contador
	pipe := m.Redis.Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Minute)
	_, err = pipe.Exec(ctx)

	if err != nil {
		return c.Next()
	}

	return c.Next()
}
