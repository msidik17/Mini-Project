package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
)

func MapBookingToBookingResponse(booking *domain.Booking) *modelsresponse.BookingResponse {
    return &modelsresponse.BookingResponse{
        ID:   booking.ID,
        Name: booking.Name,
    }
}
