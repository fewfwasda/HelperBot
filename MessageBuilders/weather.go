package messagebuilders

import (
	texts "HelperBot/Data/textsUI"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Weather(chatID int64, messageID int) tgbotapi.EditMessageTextConfig {

	degrees := 10
	weather := fmt.Sprintf(texts.WeatherText, degrees)
	weatherMessage := tgbotapi.NewEditMessageText(chatID, messageID, weather)

	tomorrow := tgbotapi.NewInlineKeyboardButtonData(texts.TomorrowButtonText, texts.TomorrowButtonData)
	nextWeek := tgbotapi.NewInlineKeyboardButtonData(texts.NextWeekButtonText, texts.NextWeekButtonData)
	back := tgbotapi.NewInlineKeyboardButtonData(texts.BackButtonText, texts.BackButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tomorrow, nextWeek),
		tgbotapi.NewInlineKeyboardRow(back),
	)

	weatherMessage.ReplyMarkup = &keyboard

	return weatherMessage
}

func WeatherMessage(chatID int64) tgbotapi.MessageConfig {
	degrees := 10
	weather := fmt.Sprintf(texts.WeatherText, degrees)
	weatherMessage := tgbotapi.NewMessage(chatID, weather)

	tomorrow := tgbotapi.NewInlineKeyboardButtonData(texts.TomorrowButtonText, texts.TomorrowButtonData)
	nextWeek := tgbotapi.NewInlineKeyboardButtonData(texts.NextWeekButtonText, texts.NextWeekButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tomorrow, nextWeek),
	)
	weatherMessage.ReplyMarkup = &keyboard

	return weatherMessage
}
