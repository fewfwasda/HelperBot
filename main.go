package main

import (
	query "HelperBot/BotFunctionality/Handler/Query"
	usertext "HelperBot/BotFunctionality/Handler/UserText"
	setcommand "HelperBot/BotFunctionality/setCommand"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env file %v", err)
	}
	botToken := os.Getenv("TOKEN_TELEGRAM_BOT")
	if botToken == "" {
		log.Fatalf("Token is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	setcommand.SetCommand(bot)

	log.Printf("Autorization on account %v", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			usertext.UserTextHandler(bot, update)
		}
		if update.CallbackQuery != nil {
			query.QueryHandler(bot, update)
		}
	}

}
