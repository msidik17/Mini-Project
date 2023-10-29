package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
)

func MapCreateBookingRequestToBooking(req *modelsrequest.CreateBookingRequest) *domain.Booking {
	return &domain.Booking{
		Name: req.Name,
	}
}
