package domain


type Movie struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Tickets     []Ticket
}
