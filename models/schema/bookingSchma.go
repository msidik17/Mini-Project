package schema

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID        uint `gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
