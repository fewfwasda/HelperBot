package initializationbot

import (
	"database/sql"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
	DB     *sql.DB
}

var Instance *Bot

func NewBot() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	botToken := os.Getenv("TOKEN_TELEGRAM_BOT")
	if botToken == "" {
		return fmt.Errorf("token is not set %v", err)
	}

	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		db.Close()
		return err
	}

	Instance = &Bot{BotAPI: botAPI, DB: db}
	return err
}
