package writer

import (
	"sahaj/flight/dto"
)

type Writer interface {
	Write(processedTickets []dto.ValidTicketCSV, failedTickets []dto.InvalidTicketCSV) error
}
