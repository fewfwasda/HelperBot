package query

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	botstatestext "HelperBot/Data/botStatesText"
	texts "HelperBot/Data/textsUI"
	messagebuilders "HelperBot/MessageBuilders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CallbackHandlers = map[string]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64){
	texts.MyNotesData:       handleMyNotes,
	texts.ShowWeatherData:   handleShowWeather,
	texts.ShowAllNotesData:  handleShowAllNotes,
	texts.ClearAllNotesData: handleClearAllNotes,
	texts.AddNoteData:       handleAddNote,
	texts.DeleteNoteData:    handleDeleteNote,
}

func handleMyNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.MyNotes(chatID, messageID))
}

func handleShowWeather(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.Weather(chatID, messageID, 15))
}

func handleShowAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	messageText := notemanager.ShowNoteList(int(userID))

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)
}

func handleClearAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	notemanager.ClearNoteList(int(userID))

	msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserClearNote)
	bot.Send(msg)
}

func handleAddNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	botstates.SetSate(int(userID), botstatestext.WaitAddNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
	bot.Send(msg)
}

func handleDeleteNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	botstates.SetSate(int(userID), botstatestext.WaitDeleteNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
	bot.Send(msg)
}
