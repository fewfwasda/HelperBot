-- -- +goose Up
-- CREATE TABLE bot_command (
--     id SERIAL,
--     title text,
--     content text,
--     PRIMARY KEY (id)
-- );

-- INSERT INTO bot_command (title, content) VALUES
--     ('ShowWeather', 'show_weather'),
--     ('AddNote', 'add_note'),
--     ('DeleteNote', 'delete_note'),
--     ('ShowAllNotes', 'show_all_notes'),
--     ('ClearAllNotes', 'clear_all_notes'),
--     ('ConfigUserNotes', 'my_notes'),

--     ('ShowWeatherDesc', 'Узнать текущую погоду в вашем городе 🔅'),
--     ('AddNoteDesc', 'Добавить новую заметку ✍🏻'),
--     ('DeleteNoteDesc', 'Удалить заметку по номеру'),
--     ('ShowAllNotesDesc', 'Показать список всех ваших заметок 📑'),
--     ('ClearAllNotesDesc', 'Удалить все заметки сразу 🧼'),
--     ('ConfigUserNotesDesc', 'Управление заметками 📝');

-- -- +goose Down
--     DROP TABLE IF EXISTS bot_command;