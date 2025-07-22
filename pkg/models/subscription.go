package models

import "time"

// модель подписки
type Subscription struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Plan      string  `gorm:"type:varchar(50);not null"`
	Price     float64 //`gorm:"not null"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
