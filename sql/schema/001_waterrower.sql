-- +goose Up
CREATE SCHEMA fitness;

CREATE TABLE fitness.waterrower (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    stroke_rate INTEGER NOT NULL,
    total_strokes INTEGER NOT NULL,
    total_distance_m INTEGER NOT NULL,
    instantaneous_pace REAL NOT NULL,
    speed INTEGER NOT NULL,
    watts INTEGER NOT NULL,
    total_kcal REAL NOT NULL,
    total_kcal_hour INTEGER NOT NULL,
    total_kcal_min INTEGER NOT NULL,
    heart_rate INTEGER NOT NULL,
    elapsedtime INTEGER NOT NULL
);

-- +goose Down
DROP TABLE fitness.waterrower;
DROP SCHEMA fitness;