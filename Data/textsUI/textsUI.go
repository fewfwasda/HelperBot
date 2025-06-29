package texts

const (
	WelcomeText = "Привет, друг! 👋 Я — твой персональный Помощник, и я здесь, чтобы сделать твой день удобнее!✨\n" +
		"Что я умею:\n☀️ Погода — Узнай текущую погоду в любом городе.\n📝 Заметки — Хочешь ничего не забыть? Добавляй, удаляй и просматривай заметки в пару кликов."
	WeatherText         = "%v - градусов в вашем городе.\nВы также можете узнать какая погода будет"
	NotesText           = "Выберите действие, которое хотите проделать с заметками"
	LocationRequestText = "Поделитесь своим местоположением"

	ConfigUserNotesButtonText = "Управление заметками 📝"
	ShowWeatherButtonText     = "Показать погоду 🔅"
	ShowAllNotesButtonText    = "Показать все заметки 📑"
	ClearAllNotesButtonText   = "Очистить все заметки 🧼"
	AddNoteButtonText         = "Добавить заметку ✍🏻"
	DeleteNoteButtonText      = "Удалить заметку по номеру"
	BackButtonText            = "Назад 🔙"
	TomorrowButtonText        = "Завтра ⛅️"
	NextWeekButtonText        = "На неделю 📅"
	LocationRequestButtonText = "📍 Поделиться своим местоположением"

	UserNotesButtonData       = "userNotes"
	ShowWeatherButtonData     = "showWeather"
	ShowAllNotesButtonData    = "showAllNotes"
	ClearAllNotesButtonData   = "clearAllNotes"
	AddNoteButtonData         = "addNote"
	DeleteNoteButtonData      = "deleteNote"
	BackButtonData            = "back"
	TomorrowButtonData        = "tomorrow"
	NextWeekButtonData        = "nextWeek"
	LocationRequestButtonData = "locationRequest"

	WaitingAddNote    = "Введите текст заметки, которую хотите добавить"
	WaitingDeleteNote = "Введите номер заметки, которую хотите удалить"

	ReplyToUserAddNote      = "Заметка '%v' успешно добавлена ✅"
	ReplyToUserDeleteNote   = "Задача '%v' успешно удалена  ✅"
	ReplyToUserClearNote    = "Все ваши заметки удалены ✅"
	ReplyToUserShowAllNotes = "Все ваши заметки:\n"
	ReplyToUserShowWeather  = "Погода в %v: %v градусов"
	ReplyToUserSaveCity     = "Ваш город сохранен ✅\nСейчас покажу погоду"

	ErrEmptyNoteList     = "Ваш список задач пуст ❌. Вы можете добавить задачу нажав на кнопку 'Добавить заметку' или при помощи команды /add_note"
	ErrFailDeleteNote    = "Невозможно удалить задачу с номером %v ❌"
	ErrUndefinedButton   = "Неизвестная кнопка ❌"
	ErrUndefinedCommand  = "Неизвесткая команда ❌"
	ErrToInstallCommands = "Не удалось установить команды: %v ❌"
	ErrSaveCity          = "Не удалось сохранить город ❌"
)
