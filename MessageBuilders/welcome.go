package messagebuilders

import (
	texts "HelperBot/Data/textsUI"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Welcome(chatID int64, messageID int) tgbotapi.EditMessageTextConfig {
	welcomeMessage := tgbotapi.NewEditMessageText(chatID, messageID, texts.WelcomeText)

	userNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ConfigUserNotesButtonText, texts.UserNotesButtonData)
	wether := tgbotapi.NewInlineKeyboardButtonData(texts.ShowWeatherButtonText, texts.ShowWeatherButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(userNotes, wether),
	)

	welcomeMessage.ReplyMarkup = &keyboard

	return welcomeMessage
}

func WelcomeMessage(chatID int64) tgbotapi.MessageConfig {
	welcomeMessage := tgbotapi.NewMessage(chatID, texts.WelcomeText)

	userNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ConfigUserNotesButtonText, texts.UserNotesButtonData)
	wether := tgbotapi.NewInlineKeyboardButtonData(texts.ShowWeatherButtonText, texts.ShowWeatherButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(userNotes, wether),
	)

	welcomeMessage.ReplyMarkup = &keyboard

	return welcomeMessage
}
