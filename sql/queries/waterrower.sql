-- name: SaveRowerData :one
INSERT INTO waterrower (id, created_at, stroke_rate, total_strokes, total_distance_m, instantaneous_pace, speed, watts, total_kcal, total_kcal_hour, total_kcal_min, heart_rate, elapsedtime)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13
)
RETURNING *;
