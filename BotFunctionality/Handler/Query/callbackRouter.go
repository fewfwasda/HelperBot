package query

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	"HelperBot/BotFunctionality/weather"
	botstatestext "HelperBot/Data/botStatesText"
	texts "HelperBot/Data/textsUI"
	initializationbot "HelperBot/InitializationBot"
	messagebuilders "HelperBot/MessageBuilders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CallbackHandlers = map[string]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64){
	texts.UserNotesButtonData:       handleConfigUserNotes,
	texts.ShowAllNotesButtonData:    handleShowAllNotes,
	texts.ClearAllNotesButtonData:   handleClearAllNotes,
	texts.AddNoteButtonData:         handleAddNote,
	texts.DeleteNoteButtonData:      handleDeleteNote,
	texts.BackButtonData:            handleBack,
	texts.ShowWeatherButtonData:     handleShowWeather,
	texts.LocationRequestButtonData: handleLocationRequest,
	texts.ChangeCityButtonText:      handleChangeCity,
}

func handleConfigUserNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.ConfigUserNotes(chatID, messageID))
}

func handleShowAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID
	db := initializationbot.Instance.DB

	messageText := notemanager.NoteList(db, userID)

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)
}

func handleClearAllNotes(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID
	db := initializationbot.Instance.DB
	notemanager.ClearNoteList(db, userID)

	msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserClearNote)
	bot.Send(msg)
}

func handleAddNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	botstates.SetState(userID, botstatestext.WaitAddNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
	bot.Send(msg)
}

func handleDeleteNote(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID

	botstates.SetState(userID, botstatestext.WaitDeleteNote)

	msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
	bot.Send(msg)
}

func handleBack(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	messageID := update.CallbackQuery.Message.MessageID
	bot.Send(messagebuilders.Welcome(chatID, messageID))
}

func handleShowWeather(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID
	db := initializationbot.Instance.DB
	messageID := update.CallbackQuery.Message.MessageID
	if city, ok := weather.GetUserCity(db, userID); ok {
		emoji, desc, temp, err := weather.GetWeather(city)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, err.Error()))
			return
		}
		bot.Send(messagebuilders.Weather(chatID, messageID, emoji, desc, city, temp))
		return
	}
	botstates.SetState(userID, botstatestext.WaitLocation)
	bot.Send(messagebuilders.LocationRequestMessage(chatID))
}

func handleChangeCity(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	userID := update.CallbackQuery.From.ID
	botstates.SetState(userID, botstatestext.WaitLocation)
	bot.Send(messagebuilders.LocationRequestMessage(chatID))
}

func handleLocationRequest(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64) {
	bot.Send(messagebuilders.LocationRequestMessage(chatID))
}
