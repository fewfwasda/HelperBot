package usertexthandler

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	"HelperBot/BotFunctionality/weather"
	botcommandtext "HelperBot/Data/botCommandText"
	botstatestext "HelperBot/Data/botStatesText"
	texts "HelperBot/Data/textsUI"
	messagebuilders "HelperBot/MessageBuilders"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UserTextHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID
	userMessage := update.Message.Text

	if botstates.GetState(userID) == botstatestext.WaitLocation {
		handleUserCity(bot, update, chatID, userID)
		return
	}
	botState := botstates.GetState(userID)

	switch botState {
	case botstatestext.WaitAddNote:
		answer := notemanager.AddNote(userID, userMessage)
		botstates.ClearState(userID)

		msg := tgbotapi.NewMessage(chatID, answer)
		bot.Send(msg)
		return
	case botstatestext.WaitDeleteNote:

		answer := notemanager.DeleteNote(userID, userMessage)
		botstates.ClearState(userID)

		msg := tgbotapi.NewMessage(chatID, answer)

		bot.Send(msg)
		return
	}

	switch userMessage {
	case "/" + botcommandtext.ConfigUserNotes:
		bot.Send(messagebuilders.ConfigUserNotesMessage(chatID))
	case "/" + botcommandtext.AddNote:
		botstates.SetState(userID, botstatestext.WaitAddNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
		bot.Send(msg)
	case "/" + botcommandtext.ClearAllNotes:
		notemanager.ClearNoteList(userID)
		msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserClearNote)
		bot.Send(msg)
	case "/" + botcommandtext.DeleteNote:
		botstates.SetState(userID, botstatestext.WaitDeleteNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
		bot.Send(msg)
	case "/" + botcommandtext.ShowAllNotes:
		messageText := notemanager.ShowNoteList(userID)
		msg := tgbotapi.NewMessage(chatID, messageText)
		bot.Send(msg)
	case "/" + botcommandtext.ShowWeather:
		if city, ok := weather.GetUserCity(userID); ok {
			emoji, desc, temp, err := weather.GetWeather(city)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(chatID, err.Error()))
				return
			}
			bot.Send(messagebuilders.Weather(chatID, emoji, desc, city, temp))
			return
		}
		botstates.SetState(userID, botstatestext.WaitLocation)
		bot.Send(messagebuilders.LocationRequestMessage(chatID))
	default:
		bot.Send(messagebuilders.WelcomeMessage(chatID))
	}
}

// Преобразовывает координаты пользователя в город и сохраняет, выводя погоду города
func handleUserCity(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64, userID int64) {
	botstates.ClearState(userID)
	// обработка кнопки
	if update.Message.Location != nil {

		coordinates := update.Message.Location

		lat := coordinates.Latitude
		lon := coordinates.Longitude

		city, err := weather.ReverseGeocode(lat, lon)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, texts.ErrCouldNotFindCity))
			return
		}

		err = weather.SetCity(userID, city)

		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, texts.ErrSaveCity))
			return
		}

		emoji, desc, temp, err := weather.GetWeather(city)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, texts.ErrGetWeather))
			return
		}

		removeKB := tgbotapi.NewRemoveKeyboard(true)
		confirm := tgbotapi.NewMessage(chatID, fmt.Sprintf(texts.ReplyToUserSaveCity))
		confirm.ReplyMarkup = removeKB
		bot.Send(confirm)

		bot.Send(messagebuilders.Weather(chatID, emoji, desc, city, temp))
		return
	}
	removeKB := tgbotapi.NewRemoveKeyboard(true)

	city := update.Message.Text
	emoji, desc, temp, err := weather.GetWeather(city)
	if err != nil {
		err := tgbotapi.NewMessage(chatID, texts.ErrCouldNotFindCity)
		err.ReplyMarkup = removeKB
		bot.Send(err)
		return
	}

	err = weather.SetCity(userID, city)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, texts.ErrSaveCity))
		return
	}
	confirm := tgbotapi.NewMessage(chatID, fmt.Sprintf(texts.ReplyToUserSaveCity))
	confirm.ReplyMarkup = removeKB
	bot.Send(confirm)

	bot.Send(messagebuilders.Weather(chatID, emoji, desc, city, temp))
}
