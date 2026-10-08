package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string
	TelegramWebhookURL string
	DatabaseURL string
	RedisURL string
	Environment string
	LogLevel string
}

var AppConfig *Config

func LoadConfig() *Config {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	viper.SetEnvPrefix("BOT")
	viper.AutomaticEnv()

	config := &Config{
		TelegramBotToken:   getEnv("TELEGRAM_BOT_TOKEN", ""),
		TelegramWebhookURL: getEnv("TELEGRAM_WEBHOOK_URL", ""),
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/botdb?sslmode=disable"),
		RedisURL:           getEnv("REDIS_URL", "localhost:6379"),
		Environment:        getEnv("ENVIRONMENT", "development"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
	}

	AppConfig = config
	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
