package domain

type Booking struct {
    ID   uint `gorm:"primaryKey"`
    Name string
    Tickets []Ticket
}