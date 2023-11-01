package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/models/schema"
)

type OrderCreateRequest struct {
	UserId     uint    `json:"user_id"`
	MovieId    uint    `json:"movie_id"`
	Quantity   int     `json:"quantity" form:"quantity" validate:"required,min=1"`
	SeatNumber string  `json:"SeatNumber" form:"SeatNumber" validate:"required,min=1,max=3"`
	TotalPrice float64 `json:"total_price" form:"total_price" validate:"required,min=1"`
}

func OrderDomaintoOrderSchema(Order domain.Order) *schema.Order {
	return &schema.Order{
		ID:         Order.ID,
		UserID:     Order.UserID,
		MovieID:    Order.MovieID,
		Quantity:   Order.Quantity,
		TotalPrice: Order.TotalPrice,
	}
}

func OrderCreateRequestToOrderDomain(request *modelsrequest.CreateOrder) *domain.Order {
	return &domain.Order{
		UserID:     request.UserID,
		MovieID:    request.MovieID,
		Quantity:   request.Quantity,
		TotalPrice: request.TotalPrice,
	}
}