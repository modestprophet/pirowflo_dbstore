-- +goose Up
ALTER TABLE fitness.waterrower
    ADD COLUMN timestamp TIMESTAMP,
    ADD COLUMN workout_id UUID;

-- +goose Down
ALTER TABLE fitness.waterrower
    DROP COLUMN timestamp,
    DROP COLUMN workout_id;