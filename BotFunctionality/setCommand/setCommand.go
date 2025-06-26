package setcommand

import (
	botcommandText "HelperBot/Data/botCommandText"
	texts "HelperBot/Data/textsUI"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCommand(bot *tgbotapi.BotAPI) error {
	commands := []tgbotapi.BotCommand{
		{
			Command:     botcommandText.AddNote,
			Description: botcommandText.AddNoteDesc,
		},
		{
			Command:     botcommandText.DeleteNote,
			Description: botcommandText.DeleteNoteDesc,
		},
		{
			Command:     botcommandText.ShowAllNotes,
			Description: botcommandText.ShowAllNotesDesc,
		},
		{
			Command:     botcommandText.ClearAllNotes,
			Description: botcommandText.ClearAllNotesDesc,
		},
		{
			Command:     botcommandText.ShowWeather,
			Description: botcommandText.ShowWeatherDesc,
		},

		{
			Command:     botcommandText.ConfigUserNotes,
			Description: botcommandText.ConfigUserNotesDesc,
		},
	}
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	_, err := bot.Request(setCommands)
	if err != nil {
		log.Printf(texts.ErrToInstallCommands, err)
		return err
	}
	return nil
}
