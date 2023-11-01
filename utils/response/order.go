package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

type OrderResponse struct {
	ID         uint    `json:"id"`
	UserId     uint    `json:"user_id"`
	MovieId    uint    `json:"movie_id"`
	SeatNumber string  `json:"seat_number"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

func ConvertSchematoResponse(order *schema.Order) *domain.Order {
	return &domain.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		MovieID:    order.MovieID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
	}
}

func OrderSchematoOrderDomain(Order *schema.Order) *domain.Order {
	return &domain.Order{
		ID:         Order.ID,
		UserID:     Order.UserID,
		MovieID:    Order.MovieID,
		Quantity:   Order.Quantity,
		TotalPrice: Order.TotalPrice,
	}
}

func CreateOrderToOrderResponse(Order *domain.Order) modelsresponse.CreateOrderResponse {
	return modelsresponse.CreateOrderResponse{
		ID:         Order.ID,
		UserID:     Order.UserID,
		MovieID:    Order.MovieID,
		Quantity:   Order.Quantity,
		TotalPrice: Order.TotalPrice,
	}
}

func ConvertOrderResponse(orders []domain.Order) []modelsresponse.OrderResponse {
	var results []modelsresponse.OrderResponse
	for _, Order := range orders {
		orderResponse := modelsresponse.OrderResponse{
			ID:         Order.ID,
			UserID:     Order.UserID,
			MovieID:    Order.MovieID,
			Quantity:   Order.Quantity,
			TotalPrice: Order.TotalPrice,
		}
		results = append(results, orderResponse)
	}
	return results
}
