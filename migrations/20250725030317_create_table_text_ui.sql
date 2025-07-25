-- +goose Up
CREATE TABLE text_ui (
    id serial,
    title text,
    content text,
    primary key (id)
);

INSERT INTO text_ui (title, content) VALUES
    ('WelcomeText', 'Привет, друг! 👋 Я — твой персональный Помощник, и я здесь, чтобы сделать твой день удобнее!✨\nЧто я умею:\n☀️ Погода — Узнай текущую погоду в любом городе.\n📝 Заметки — Хочешь ничего не забыть? Добавляй, удаляй и просматривай заметки в пару кликов.'),
    ('ConfigUserNotesButtonText', 'Управление заметками 📝'),
    ('ShowWeatherButtonText', 'Показать погоду 🔅'),
    ('ShowAllNotesButtonText', 'Показать все заметки 📑'),
    ('ClearAllNotesButtonText', 'Очистить все заметки 🧼'),
    ('AddNoteButtonText', 'Добавить заметку ✍🏻'),
    ('DeleteNoteButtonText', 'Удалить заметку по номеру'),
    ('BackButtonText', 'Назад 🔙'),
    ('ChangeCityButtonText', 'Сменить город 🏙'),
    ('LocationRequestButtonText', 'Поделиться своим местоположением 📍'),
    ('UserNotesButtonData', 'userNotes'),
    ('ShowWeatherButtonData', 'showWeather'),
    ('ShowAllNotesButtonData', 'showAllNotes'),
    ('ClearAllNotesButtonData', 'clearAllNotes'),
    ('AddNoteButtonData', 'addNote'),
    ('DeleteNoteButtonData', 'deleteNote'),
    ('BackButtonData', 'back'),
    ('ChangeCityButtonData', 'changeCity'),
    ('LocationRequestButtonData', 'locationRequest'),
    ('WaitingAddNote', 'Введите текст заметки, которую хотите добавить'),
    ('WaitingDeleteNote', 'Введите номер заметки, которую хотите удалить'),
    ('ReplyToUserNotesText', 'Выберите действие, которое хотите проделать с заметками'),
    ('ReplyToUserAddNote', 'Заметка %v успешно добавлена ✅'),
    ('ReplyToUserDeleteNote', 'Заметка под номером %v успешно удалена  ✅'),
    ('ReplyToUserClearNote', 'Все ваши заметки удалены ✅'),
    ('ReplyToUserShowAllNotes', 'Все ваши заметки:\n'),
    ('ReplyToUserShowWeather', '%v - градусов в городе %v.\n%v %v'),
    ('ReplyToUserSaveCity', 'Ваш город сохранен ✅\nСейчас покажу погоду'),
    ('ReplyToUserLocationRequestText', 'Нажмите на кнопку \"📍 Поделиться своим местоположением\"\nИли введите название города самостоятельно'),
    ('ErrEmptyNoteList', 'у вас нет сохраненных заметок ❌. вы можете добавить заметку нажав на кнопку "добавить заметку" или при помощи команды /add_note'),
    ('ErrDeleteNote', 'невозможно удалить заметку под номером номером %v ❌'),
    ('ErrUndefinedButton', 'неизвестная кнопка ❌'),
    ('ErrUndefinedCommand', 'неизвесткая команда ❌'),
    ('ErrToInstallCommands', 'не удалось установить команды: %v ❌'),
    ('ErrSaveCity', 'не удалось сохранить город ❌'),
    ('ErrGetWeather', 'не удалось получить погоду ❌'),
    ('ErrAPIKey', 'api ключ не установлен ❌'),
    ('ErrCouldNotFindCity', 'не удалось найти город ❌'),
    ('ErrAddNote', 'не удалось схранить заметку ❌');

-- +goose Down
    DROP TABLE IF EXISTS text_ui;