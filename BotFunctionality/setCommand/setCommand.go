package setcommand

import (
	texts "HelperBot/Data/textsUI"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCommand(bot *tgbotapi.BotAPI) error {
	commands := []tgbotapi.BotCommand{
		{
			Command:     "showWeather",
			Description: texts.ShowWeatherButtonText,
		},
		{
			Command:     "clearNotes",
			Description: texts.ShowWeatherButtonText,
		},
		{
			Command:     "showNotes",
			Description: texts.ShowAllNotesButtonText,
		},
	}
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	_, err := bot.Request(setCommands)
	if err != nil {
		log.Printf("Не удалось установить команды: %v", err)
		return err
	}
	return nil
}
