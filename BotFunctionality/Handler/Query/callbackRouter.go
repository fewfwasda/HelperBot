package query

import (
	"HelperBot/Data/texts"
	messagebuilders "HelperBot/MessageBuilders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CallbackHandlers = map[string]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64){
	texts.MyNotesData:     handleMyNotes,
	texts.ShowWeatherData: handleShowWeather,
	// texts.ShowAllNotesData:  handleShowAllNotes,
	// texts.ClearAllNotesData: handleClearAllNotes,
	// texts.AddNoteData:       handleAddNote,
	// texts.DeleteNoteData:    handleDeleteNote,
}

func handleMyNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	bot.Send(messagebuilders.MyNotes(chatID))
}

func handleShowWeather(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	bot.Send(messagebuilders.Weather(chatID, 15))
}

// func handleShowAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
// 	answer := tgbotapi.NewMessage(chatID, "ggg")
// 	bot.Send(answer)
// }

// func handleClearAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
// 	bot.Send(messagebuilders.MyNotes(chatID))
// }

// func handleAddNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
// 	bot.Send(messagebuilders.MyNotes(chatID))
// }
// func handleDeleteNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
// 	bot.Send(messagebuilders.MyNotes(chatID))
// }
