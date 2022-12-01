package reader

import (
	"sahaj/flight/dto"
)

type Reader interface {
	Read() ([]dto.TicketCSV, error)
}
