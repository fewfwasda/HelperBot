package usertext

import (
	botstates "HelperBot/BotFunctionality/botStates"
	notemanager "HelperBot/BotFunctionality/noteManager"
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
		notemanager.AddNote(int(userID), userMessage)
		botstates.ClearState(int(userID))

		msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserAddNote)
		bot.Send(msg)
		return
	case botstatestext.WaitDeleteNote:
		notemanager.DeleteNote(int(userID), userMessage)
		botstates.ClearState(int(userID))

		msg := tgbotapi.NewMessage(chatID, texts.ReplyToUserDeleteNote)
		bot.Send(msg)
		return
	}

	switch userMessage {
	case "/start":
		bot.Send(messagebuilders.Welcome(chatID))
	default:
		unknownCommand := tgbotapi.NewMessage(chatID, "Неизвесткая команда(")
		bot.Send(unknownCommand)
	}
}
