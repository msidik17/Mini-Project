package modelsresponse

type OrderResponse struct {
	ID         uint    `json:"id"`
	UserID     uint    `json:"user_id"`
	MovieID    uint    `json:"movie_id"`
	SeatNumber string  `json:"seat_number"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CreateOrderResponse struct {
	ID         uint    `json:"id"`
	UserID     uint    `json:"user_id"`
	MovieID    uint    `json:"movie_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}