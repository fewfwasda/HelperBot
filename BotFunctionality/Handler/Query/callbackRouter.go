package query

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	weathermanager "HelperBot/BotFunctionality/weatherManager"
	botstatestext "HelperBot/Data/botStatesText"
	texts "HelperBot/Data/textsUI"
	messagebuilders "HelperBot/MessageBuilders"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CallbackHandlers = map[string]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64){
	texts.UserNotesButtonData:       handleConfigUserNotes,
	texts.ShowWeatherButtonData:     handleShowWeather,
	texts.ShowAllNotesButtonData:    handleShowAllNotes,
	texts.ClearAllNotesButtonData:   handleClearAllNotes,
	texts.AddNoteButtonData:         handleAddNote,
	texts.DeleteNoteButtonData:      handleDeleteNote,
	texts.BackButtonData:            handleBack,
	texts.LocationRequestButtonData: handleLocationRequest,
}

func handleConfigUserNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.ConfigUserNotes(chatID, messageID))
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

	botstates.SetState(int(userID), botstatestext.WaitAddNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
	bot.Send(msg)
}

func handleDeleteNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	botstates.SetState(int(userID), botstatestext.WaitDeleteNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
	bot.Send(msg)
}

func handleBack(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.Welcome(chatID, messageID))
}

func handleShowWeather(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	if city, ok := weathermanager.GetUserCity(int(userID)); ok {
		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf(texts.ReplyToUserShowWeather, city, 15)))
		return
	}
	botstates.SetState(int(userID), botstatestext.WaitLocation)
	bot.Send(messagebuilders.LocationRequestMessage(chatID))
}

func handleLocationRequest(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	bot.Send(messagebuilders.LocationRequestMessage(chatID))
}
