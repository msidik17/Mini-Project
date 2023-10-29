package schema

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID        uint           `gorm:"primaryKey"`
	MovieID   uint           `gorm:"index"`
	Movie     Movie          `gorm:"foreignKey:MovieID"`
	BookingID uint           `gorm:"index"`
	Booking   Booking        `gorm:"foreignKey:BookingID"`
	StudioID  uint           `gorm:"index"`
	Studio    Studio         `gorm:"foreignKey:StudioID"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
