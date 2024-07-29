package models

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Workouts []Workout `gorm:"foreignKey:UserID"`
}
