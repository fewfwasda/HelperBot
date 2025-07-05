package texts

const (
	WelcomeText = "Привет, друг! 👋 Я — твой персональный Помощник, и я здесь, чтобы сделать твой день удобнее!✨\n" +
		"Что я умею:\n☀️ Погода — Узнай текущую погоду в любом городе.\n📝 Заметки — Хочешь ничего не забыть? Добавляй, удаляй и просматривай заметки в пару кликов."

	ConfigUserNotesButtonText = "Управление заметками 📝"
	ShowWeatherButtonText     = "Показать погоду 🔅"
	ShowAllNotesButtonText    = "Показать все заметки 📑"
	ClearAllNotesButtonText   = "Очистить все заметки 🧼"
	AddNoteButtonText         = "Добавить заметку ✍🏻"
	DeleteNoteButtonText      = "Удалить заметку по номеру"
	BackButtonText            = "Назад 🔙"
	ChangeCityButtonText      = "Сменить город 🏙️"
	LocationRequestButtonText = "Поделиться своим местоположением 📍"

	UserNotesButtonData       = "userNotes"
	ShowWeatherButtonData     = "showWeather"
	ShowAllNotesButtonData    = "showAllNotes"
	ClearAllNotesButtonData   = "clearAllNotes"
	AddNoteButtonData         = "addNote"
	DeleteNoteButtonData      = "deleteNote"
	BackButtonData            = "back"
	ChangeCityButtonData      = "changecity"
	LocationRequestButtonData = "locationRequest"

	WaitingAddNote    = "Введите текст заметки, которую хотите добавить"
	WaitingDeleteNote = "Введите номер заметки, которую хотите удалить"

	ReplyToUserNotesText           = "Выберите действие, которое хотите проделать с заметками"
	ReplyToUserAddNote             = "Заметка '%v' успешно добавлена ✅"
	ReplyToUserDeleteNote          = "Заметка под номером %v успешно удалена  ✅"
	ReplyToUserClearNote           = "Все ваши заметки удалены ✅"
	ReplyToUserShowAllNotes        = "Все ваши заметки:\n"
	ReplyToUserShowWeather         = "%v - градусов в городе %v.\n%v %v"
	ReplyToUserSaveCity            = "Ваш город сохранен ✅\nСейчас покажу погоду"
	ReplyToUserLocationRequestText = "Нажмите на кнопку \"📍 Поделиться своим местоположением\"\nИли введите название города самостоятельно"

	ErrEmptyNoteList     = "У вас нет сохраненных заметок ❌. Вы можете добавить заметку нажав на кнопку 'Добавить заметку' или при помощи команды /add_note"
	ErrDeleteNote        = "Невозможно удалить заметку под номером номером %v ❌"
	ErrUndefinedButton   = "Неизвестная кнопка ❌"
	ErrUndefinedCommand  = "Неизвесткая команда ❌"
	ErrToInstallCommands = "Не удалось установить команды: %v ❌"
	ErrSaveCity          = "Не удалось сохранить город ❌"
	ErrGetWeather        = "Не удалось получить погоду ❌"
	ErrAPIKey            = "API ключ не установлен ❌"
	ErrCouldNotFindCity  = "не удалось найти город ❌"
	ErrAddNote           = "Не удалось схранить заметку ❌"
)
