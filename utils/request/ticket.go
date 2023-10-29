package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
)

func MapCreateTicketRequestToTicket(req *modelsrequest.CreateTicketRequest) *domain.Ticket {
	return &domain.Ticket{
		MovieID:   req.MovieID,
		BookingID: req.BookingID,
		StudioID:  req.StudioID,
	}
}
