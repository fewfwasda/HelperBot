-- -- +goose Up
-- CREATE TABLE bot_states_text (
--     id SERIAL,
--     title text,
--     content text,
--     PRIMARY KEY (id)
-- );

-- INSERT INTO bot_states_text (title, content) VALUES
--     ('WaitAddNote', 'wait_add_note'),
--     ('WaitDeleteNote', 'wait_delete_note'),
--     ('WaitLocation', 'wait_location');

-- -- +goose Down
--     DROP TABLE IF EXISTS bot_states_text;