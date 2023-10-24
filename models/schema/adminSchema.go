package schema

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"PrimaryKey"`
	Name      string         `gorm:"Name"`
	Email     string         `gorm:"Email"`
	Password  string         `gorm:"Password"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
