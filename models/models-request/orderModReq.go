package modelsrequest

type OrderCreateRequest struct {
	UserId     uint    `json:"user_id"`
	MovieId    uint    `json:"movie_id"`
	Quantity   int     `json:"quantity" form:"quantity" validate:"required,min=1"`
	SeatNumber string  `json:"SeatNumber" form:"SeatNumber" validate:"required,min=1,max=3"`
	TotalPrice float64 `json:"total_price" form:"total_price" validate:"required,min=1"`
}

type CreateOrder struct {
	UserID     uint    `json:"user_id"`
	MovieID    uint    `json:"movie_id"`
	Quantity   int     `json:"quantity" form:"quantity" validate:"required,min=1"`
	TotalPrice float64 `json:"totalPrice" form:"totalPrice" validate:"required,min=1"`
}
