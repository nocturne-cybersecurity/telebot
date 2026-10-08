package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"telegram-bot/bot"
	"telegram-bot/config"
	"telegram-bot/database"
	"telegram-bot/utils"

	"go.uber.org/zap"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar logger
	if err := utils.InitLogger(cfg.LogLevel); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer utils.SyncLogger()

	logger := utils.GetLogger()
	logger.Info("Starting Telegram Bot",
		zap.String("environment", cfg.Environment),
	)

	// Inicializar base de datos
	if err := database.InitDB(cfg); err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer database.CloseDB()

	// Inicializar Redis
	if err := database.InitRedis(cfg); err != nil {
		logger.Warn("Failed to initialize Redis, continuing without cache", zap.Error(err))
	} else {
		defer database.CloseRedis()
	}

	// Inicializar bot
	botInstance, err := bot.NewBot(cfg, database.GetDB())
	if err != nil {
		logger.Fatal("Failed to initialize bot", zap.Error(err))
	}

	// Configurar handlers
	botInstance.SetupHandlers(database.GetDB())

	// Manejar señales de terminación
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar bot en una goroutine
	go func() {
		botInstance.Start()
	}()

	// Esperar señal de terminación
	<-sigChan
	logger.Info("Shutting down bot...")
	botInstance.Stop()
	logger.Info("Bot stopped successfully")
}
