package models

import "time"

type ExerciseLog struct {
	ID         uint `gorm:"primaryKey"`
	WorkoutID  uint
	ExerciseID uint
	Reps       int
	Weight     float64
	Time       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
