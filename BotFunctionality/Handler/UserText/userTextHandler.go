package usertext

import (
	messagebuilders "HelperBot/MessageBuilders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UserTextHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userMessage := update.Message.Text

	switch userMessage {
	case "/start":
		bot.Send(messagebuilders.Welcome(chatID))
	default:
		unknownCommand := tgbotapi.NewMessage(chatID, "Неизвесткая команда(")
		bot.Send(unknownCommand)
	}
}
