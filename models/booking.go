package models

import "time"

type Booking struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	GymID     uint      `gorm:"not null" json:"gym_id"`
	Date      time.Time `json:"date"`        // День брони
	StartTime time.Time `json:"start_time"`  // Начало
	EndTime   time.Time `json:"end_time"`    // Конец
	CreatedAt time.Time `json:"created_at"`
}
