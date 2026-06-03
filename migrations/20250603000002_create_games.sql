-- +goose Up
CREATE TABLE games (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name      TEXT NOT NULL,
    publisher TEXT NOT NULL,
    year      TEXT NOT NULL,
    platform  TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS games;
