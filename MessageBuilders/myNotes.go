package messagebuilders

import (
	texts "HelperBot/Data/textsUI"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConfigUserNotes(chatID int64, messageID int) tgbotapi.EditMessageTextConfig {
	userNotesMessage := tgbotapi.NewEditMessageText(chatID, messageID, texts.NotesText)

	showAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ShowAllNotesButtonText, texts.ShowAllNotesButtonData)
	clearAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ClearAllNotesButtonText, texts.ClearAllNotesButtonData)
	addNote := tgbotapi.NewInlineKeyboardButtonData(texts.AddNoteButtonText, texts.AddNoteButtonData)
	deleteNote := tgbotapi.NewInlineKeyboardButtonData(texts.DeleteNoteButtonText, texts.DeleteNoteButtonData)
	back := tgbotapi.NewInlineKeyboardButtonData(texts.BackButtonText, texts.BackButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(showAllNotes, clearAllNotes),
		tgbotapi.NewInlineKeyboardRow(addNote, deleteNote),
		tgbotapi.NewInlineKeyboardRow(back),
	)

	userNotesMessage.ReplyMarkup = &keyboard

	return userNotesMessage
}

func ConfigUserNotesMessage(chatID int64) tgbotapi.MessageConfig {
	userNotesMessage := tgbotapi.NewMessage(chatID, texts.NotesText)

	showAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ShowAllNotesButtonText, texts.ShowAllNotesButtonData)
	clearAllNotes := tgbotapi.NewInlineKeyboardButtonData(texts.ClearAllNotesButtonText, texts.ClearAllNotesButtonData)
	addNote := tgbotapi.NewInlineKeyboardButtonData(texts.AddNoteButtonText, texts.AddNoteButtonData)
	deleteNote := tgbotapi.NewInlineKeyboardButtonData(texts.DeleteNoteButtonText, texts.DeleteNoteButtonData)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(showAllNotes, clearAllNotes),
		tgbotapi.NewInlineKeyboardRow(addNote, deleteNote),
	)

	userNotesMessage.ReplyMarkup = &keyboard

	return userNotesMessage
}
