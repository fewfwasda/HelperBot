package query

import (
	texts "HelperBot/Data/textsUI"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func QueryHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	bot.Request(callback)

	chatID := update.CallbackQuery.Message.Chat.ID
	cq := update.CallbackQuery
	dataQuery := cq.Data

	handler, ok := CallbackHandlers[dataQuery]
	if !ok {
		bot.Send(tgbotapi.NewMessage(chatID, texts.ErrUndefinedButton))
		return
	}
	handler(bot, update, chatID)
}
