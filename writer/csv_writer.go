package writer

import (
	"os"
	"sahaj/flight/dto"

	"github.com/gocarina/gocsv"
)

type csvWriter struct {
	validTicketsFilePath   string
	invalidTicketsFilePath string
}

func NewCSVWriter(validTicketsFilePath, invalidTicketsFilePath string) Writer {
	return csvWriter{validTicketsFilePath: validTicketsFilePath,
		invalidTicketsFilePath: invalidTicketsFilePath}
}

func (csv csvWriter) WriteValidTickets(tickets []dto.ValidTicketCSV) error {
	file, err := os.Create(csv.validTicketsFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.MarshalFile(&tickets, file); err != nil {
		return err
	}
	return nil
}

func (csv csvWriter) WriteInvalidTickets(tickets []dto.InvalidTicketCSV) error {
	file, err := os.Create(csv.invalidTicketsFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := gocsv.MarshalFile(&tickets, file); err != nil {
		return err
	}
	return nil
}
