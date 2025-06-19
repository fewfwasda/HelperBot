package messagebuilders

import (
	texts "HelperBot/Data/textsUI"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Weather(chatID int64, messageID int, degrees int) tgbotapi.EditMessageTextConfig {
	weatherMessage := tgbotapi.NewEditMessageText(chatID, messageID, strconv.Itoa(degrees)+texts.WeatherText)

	return weatherMessage
}
