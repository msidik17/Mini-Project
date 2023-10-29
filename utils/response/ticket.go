package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
)

func MapTicketToTicketResponse(ticket *domain.Ticket) *modelsresponse.TicketResponse {
	return &modelsresponse.TicketResponse{
		ID:        ticket.ID,
		MovieID:   ticket.MovieID,
		BookingID: ticket.BookingID,
		StudioID:  ticket.StudioID,
	}
}
