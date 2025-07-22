-- -- +goose Up
-- CREATE TABLE emoji (
--     id SERIAL,
--     title text,
--     content text,
--     PRIMARY KEY (id)
-- );

-- INSERT INTO emoji (title, content) VALUES
--     ('Clear', '☀️'),
--     ('Clouds', '☁️'),
--     ('Rain', '🌧️'),
--     ('Snow', '❄️');

-- -- +goose Down
--     DROP TABLE IF EXISTS emoji;