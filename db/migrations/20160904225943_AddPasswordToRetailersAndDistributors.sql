
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE distributors ADD COLUMN password VARCHAR(200) UNIQUE NOT NULL;
ALTER TABLE retailers ADD COLUMN password VARCHAR(200) UNIQUE NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE distributors DROP COLUMN password;
ALTER TABLE retailers DROP COLUMN password;
