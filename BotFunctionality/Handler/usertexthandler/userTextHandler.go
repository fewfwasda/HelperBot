package usertexthandler

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	weathermanager "HelperBot/BotFunctionality/weatherManager"
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

	if update.Message.Location != nil {
		handleUserCity(bot, update, chatID, userID)
		return
	}

	botState := botstates.GetState(int(userID))

	switch botState {
	case botstatestext.WaitAddNote:
		answer := notemanager.AddNote(int(userID), userMessage)
		botstates.ClearState(int(userID))

		msg := tgbotapi.NewMessage(chatID, answer)
		bot.Send(msg)
		return
	case botstatestext.WaitDeleteNote:

		answer := notemanager.DeleteNote(int(userID), userMessage)
		botstates.ClearState(int(userID))

		msg := tgbotapi.NewMessage(chatID, answer)

		bot.Send(msg)
		return
	}

	switch userMessage {
	case "/" + botcommandtext.ShowWeather:
		bot.Send(messagebuilders.WeatherMessage(chatID))
	case "/" + botcommandtext.ConfigUserNotes:
		bot.Send(messagebuilders.ConfigUserNotesMessage(chatID))
	case "/" + botcommandtext.AddNote:
		botstates.SetState(int(userID), botstatestext.WaitAddNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
		bot.Send(msg)
	case "/" + botcommandtext.ClearAllNotes:
		notemanager.ClearNoteList(int(userID))
		msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserClearNote)
		bot.Send(msg)
	case "/" + botcommandtext.DeleteNote:
		botstates.SetState(int(userID), botstatestext.WaitDeleteNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
		bot.Send(msg)
	case "/" + botcommandtext.ShowAllNotes:
		messageText := notemanager.ShowNoteList(int(userID))
		msg := tgbotapi.NewMessage(chatID, messageText)
		bot.Send(msg)
	default:
		bot.Send(messagebuilders.WelcomeMessage(chatID))
	}
}

func handleUserCity(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatID int64, userID int64) {
	if loc := update.Message.Location; loc != nil && botstates.GetState(int(userID)) == botstatestext.WaitLocation {
		botstates.ClearState(int(userID))
		city := fmt.Sprintf("%.6f,%.6f", loc.Latitude, loc.Latitude)

		if err := weathermanager.SetCity(int(userID), city); err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, texts.ErrSaveCity+err.Error()))
			return
		}
		removeKB := tgbotapi.NewRemoveKeyboard(true)
		confirm := tgbotapi.NewMessage(chatID, fmt.Sprintf(texts.ReplyToUserSaveCity))
		confirm.ReplyMarkup = removeKB
		bot.Send(confirm)

		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf(texts.ReplyToUserShowWeather, city, 15)))
	}
}
