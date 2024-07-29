package models

import "time"

type Workout struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	Date         time.Time
	ExerciseLogs []ExerciseLog `gorm:"foreignKey:WorkoutID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
