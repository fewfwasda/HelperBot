package messagebuilders

import (
	"HelperBot/Data/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Welcome(chatID int64) tgbotapi.MessageConfig {
	welcomeMessage := tgbotapi.NewMessage(chatID, texts.WelcomeText)

	myNotes := tgbotapi.NewInlineKeyboardButtonData(texts.MyNotesButtonText, texts.MyNotesData)
	wether := tgbotapi.NewInlineKeyboardButtonData(texts.ShowWeatherButtonText, texts.ShowWeatherData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(myNotes, wether),
	)

	welcomeMessage.ReplyMarkup = keyboard

	return welcomeMessage
}
