
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE distributors ADD COLUMN username VARCHAR(200) UNIQUE NOT NULL;
ALTER TABLE retailers ADD COLUMN username VARCHAR(200) UNIQUE NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE distributors DROP COLUMN username;
ALTER TABLE retailers DROP COLUMN username;
