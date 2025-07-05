package main

import (
	query "HelperBot/BotFunctionality/Handler/Query"
	"HelperBot/BotFunctionality/Handler/usertexthandler"
	setcommand "HelperBot/BotFunctionality/setCommand"
	initializationbot "HelperBot/InitializationBot"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := initializationbot.NewBot()
	if err != nil {
		log.Fatalf("Initialization failed: %v", err)
	}
	defer initializationbot.Instance.DB.Close()

	setcommand.SetCommand(initializationbot.Instance.BotAPI)

	log.Printf("Autorization on account %v", initializationbot.Instance.BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := initializationbot.Instance.BotAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			usertexthandler.UserTextHandler(initializationbot.Instance.BotAPI, initializationbot.Instance.DB, update)
		}
		if update.CallbackQuery != nil {
			query.QueryHandler(initializationbot.Instance.BotAPI, update)
		}
	}
}
