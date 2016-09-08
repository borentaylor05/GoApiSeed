
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE distributors ADD COLUMN refresh_token VARCHAR(200);
ALTER TABLE retailers ADD COLUMN refresh_token VARCHAR(200);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE distributors DROP COLUMN refresh_token;
ALTER TABLE retailers DROP COLUMN refresh_token;
