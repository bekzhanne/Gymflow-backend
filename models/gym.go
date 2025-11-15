package models

import "time"

type Gym struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:150;not null" json:"name"`
	Address     string    `gorm:"size:255;not null" json:"address"`
	Description string    `gorm:"type:text" json:"description"`
	Capacity    int       `gorm:"not null" json:"capacity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
