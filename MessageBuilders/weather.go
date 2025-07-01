package messagebuilders

import (
	texts "HelperBot/Data/textsUI"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Weather(chatID int64, emoji string, desc string, city string, temp float64) tgbotapi.MessageConfig {
	weather := fmt.Sprintf(texts.ReplyToUserShowWeather, temp, city, desc, emoji)
	weatherMessage := tgbotapi.NewMessage(chatID, weather)

	changeCity := tgbotapi.NewInlineKeyboardButtonData(texts.ChangeCityButtonText, texts.ChangeCityButtonText)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(changeCity),
	)
	weatherMessage.ReplyMarkup = &keyboard

	return weatherMessage
}
