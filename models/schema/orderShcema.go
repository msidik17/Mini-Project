package schema

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"PrimaryKey"`
	MovieID    uint           `json:"movie_id"`
	UserID     uint           `json:"user_id"`
	Quantity   int            `json:"quantity"`
	SeatNumber string         `json:"seat_number"`
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdate"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

