package writer

import (
	"sahaj/flight/dto"
)

type Writer interface {
	WriteValidTickets(tickets []dto.ValidTicketCSV) error
	WriteInvalidTickets(tickets []dto.InvalidTicketCSV) error
}
