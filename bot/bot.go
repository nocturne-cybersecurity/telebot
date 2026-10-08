package bot

import (
	"log"
	"time"
	"telegram-bot/config"
	"telegram-bot/handlers"
	"telegram-bot/middleware"
	"telegram-bot/models"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type Bot struct {
	Telebot *telebot.Bot
	Config  *config.Config
}

func NewBot(cfg *config.Config, db *gorm.DB) (*Bot, error) {
	pref := telebot.Settings{
		Token:  cfg.TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	if cfg.TelegramWebhookURL != "" {
		pref.Poller = &telebot.Webhook{
			Endpoint: &telebot.WebhookEndpoint{
				PublicURL: cfg.TelegramWebhookURL,
			},
		}
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Telebot: b,
		Config:  cfg,
	}, nil
}

func (b *Bot) SetupHandlers(db *gorm.DB) {
	// Inicializar handlers
	commandHandler := handlers.NewCommandHandler(db)
	messageHandler := handlers.NewMessageHandler(db)

	// Middleware
	loggingMiddleware := middleware.NewLoggingMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(db)

	// Comandos
	b.Telebot.Handle(telebot.OnCommandStart, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return commandHandler.HandleStart(c)
	})

	b.Telebot.Handle(telebot.OnCommandHelp, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return commandHandler.HandleHelp(c)
	})

	b.Telebot.Handle(telebot.OnCommandAbout, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return commandHandler.HandleAbout(c)
	})

	b.Telebot.Handle(telebot.OnCommandSettings, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		if err := authMiddleware(c); err != nil {
			return err
		}
		return commandHandler.HandleSettings(c)
	})

	b.Telebot.Handle(telebot.OnCommandStats, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		if err := authMiddleware(c); err != nil {
			return err
		}
		return commandHandler.HandleStats(c)
	})

	// Mensajes de texto
	b.Telebot.Handle(telebot.OnText, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return messageHandler.HandleTextMessage(c)
	})

	// Fotos
	b.Telebot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return messageHandler.HandlePhoto(c)
	})

	// Documentos
	b.Telebot.Handle(telebot.OnDocument, func(c telebot.Context) error {
		if err := loggingMiddleware(c); err != nil {
			return err
		}
		return messageHandler.HandleDocument(c)
	})

	// Callbacks
	b.Telebot.Handle(telebot.OnCallback, func(c telebot.Context) error {
		return messageHandler.HandleCallback(c)
	})
}

func (b *Bot) Start() {
	log.Println("Bot started successfully")
	b.Telebot.Start()
}

func (b *Bot) Stop() {
	log.Println("Bot stopped")
}
