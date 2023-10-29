package domain

type Studio struct {
    ID   uint `gorm:"primaryKey"`
    Name string
    Tickets []Ticket
}