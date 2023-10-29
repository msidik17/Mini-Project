package schema

import "gorm.io/gorm"
import "time"

type Movie struct {
    ID          uint `gorm:"primaryKey"`
    Title       string `json:"title"`
    Description string `json:"description"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdate"`
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}