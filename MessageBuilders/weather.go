package messagebuilders

import (
	"HelperBot/Data/texts"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Weather(chatID int64, degrees int) tgbotapi.MessageConfig {
	weatherMessage := tgbotapi.NewMessage(chatID, strconv.Itoa(degrees)+texts.WeatherText)

	return weatherMessage
}
