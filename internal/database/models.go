// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type FitnessWaterrower struct {
	ID                uuid.UUID
	CreatedAt         time.Time
	StrokeRate        int32
	TotalStrokes      int32
	TotalDistanceM    int32
	InstantaneousPace float32
	Speed             int32
	Watts             int32
	TotalKcal         float32
	TotalKcalHour     int32
	TotalKcalMin      int32
	HeartRate         int32
	Elapsedtime       int32
}
