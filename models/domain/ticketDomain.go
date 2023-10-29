package domain

type Ticket struct {
    ID        uint `gorm:"primaryKey"`
    MovieID   uint
    Movie     Movie
    BookingID uint
    Booking   Booking
    StudioID  uint
    Studio    Studio
}