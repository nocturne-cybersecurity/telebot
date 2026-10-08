package database

import (
	"context"
	"fmt"
	"log"
	"telegram-bot/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis(cfg *config.Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})

	// Test connection
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Redis connected successfully")
	return nil
}

func GetRedis() *redis.Client {
	return RedisClient
}

func CloseRedis() error {
	return RedisClient.Close()
}

func SetCache(key string, value interface{}, expiration int) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

func GetCache(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func DeleteCache(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
