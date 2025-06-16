package messagebuilders

import (
	"HelperBot/Data/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MyNotes(chatID int64) tgbotapi.MessageConfig {
	myNotesMessage := tgbotapi.NewMessage(chatID, texts.NotesText)

	showAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ShowAllNotesButtonText, texts.ShowAllNotesData)
	clearAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ClearAllNotesButtonText, texts.ClearAllNotesData)
	addNote := tgbotapi.NewInlineKeyboardButtonData(texts.AddNoteButtonText, texts.AddNoteData)
	deleteNote := tgbotapi.NewInlineKeyboardButtonData(texts.DeleteNoteButtonText, texts.DeleteNoteData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(showAllNotes, clearAllNotes),
		tgbotapi.NewInlineKeyboardRow(addNote, deleteNote),
	)

	myNotesMessage.ReplyMarkup = keyboard

	return myNotesMessage
}
