package setcommand

import (
	botcommandtext "HelperBot/Data/botCommandText"
	texts "HelperBot/Data/textsUI"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCommand(bot *tgbotapi.BotAPI) error {
	commands := []tgbotapi.BotCommand{
		{
			Command:     botcommandtext.AddNote,
			Description: botcommandtext.AddNoteDesc,
		},
		{
			Command:     botcommandtext.DeleteNote,
			Description: botcommandtext.DeleteNoteDesc,
		},
		{
			Command:     botcommandtext.ShowAllNotes,
			Description: botcommandtext.ShowAllNotesDesc,
		},
		{
			Command:     botcommandtext.ClearAllNotes,
			Description: botcommandtext.ClearAllNotesDesc,
		},
		{
			Command:     botcommandtext.ShowWeather,
			Description: botcommandtext.ShowWeatherDesc,
		},

		{
			Command:     botcommandtext.ConfigUserNotes,
			Description: botcommandtext.ConfigUserNotesDesc,
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
