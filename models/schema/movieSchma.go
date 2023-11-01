package schema

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Studio      string         `json:"studio"`
	Price       float64        `gorm:"type:int" json:"price"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdate"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
