package domain

type Order struct {
	ID         uint `gorm:"primaryKey"`
	MovieID    uint 
	UserID     uint
	Quantity   int
	TotalPrice float64

}
