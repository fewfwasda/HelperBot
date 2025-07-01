package messagebuilders

import (
	texts "HelperBot/Data/textsUI"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LocationRequestMessage(chatID int64) tgbotapi.MessageConfig {
	requestMessage := tgbotapi.NewMessage(chatID, texts.ReplyToUserLocationRequestText)
	requestButton := tgbotapi.KeyboardButton{
		Text:            texts.LocationRequestButtonText,
		RequestLocation: true,
	}
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(requestButton),
	)
	requestMessage.ReplyMarkup = keyboard
	return requestMessage
}
