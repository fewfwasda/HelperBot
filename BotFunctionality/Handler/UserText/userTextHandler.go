package usertext

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
	botcommandText "HelperBot/Data/botCommandText"
	botstatestext "HelperBot/Data/botStatesText"
	texts "HelperBot/Data/textsUI"
	messagebuilders "HelperBot/MessageBuilders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UserTextHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID
	userMessage := update.Message.Text

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
	case "/" + botcommandText.ShowWeather:
		messageID := update.CallbackQuery.Message.MessageID
		bot.Send(messagebuilders.Weather(chatID, messageID, 15))
	case "/" + botcommandText.AddNote:
		botstates.SetSate(int(userID), botstatestext.WaitAddNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingAddNote)
		bot.Send(msg)
	case "/" + botcommandText.ClearAllNotes:
		notemanager.ClearNoteList(int(userID))
		msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserClearNote)
		bot.Send(msg)
	case "/" + botcommandText.DeleteNote:
		botstates.SetSate(int(userID), botstatestext.WaitDeleteNote)
		msg := tgbotapi.NewMessage(chatID, texts.WaitingDeleteNote)
		bot.Send(msg)
	case "/" + botcommandText.ShowAllNotes:
		messageText := notemanager.ShowNoteList(int(userID))
		msg := tgbotapi.NewMessage(chatID, messageText)
		bot.Send(msg)
	default:
		bot.Send(messagebuilders.Welcome(chatID))
	}
}
