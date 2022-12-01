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

func (csv csvWriter) Write(processedTickets []dto.ValidTicketCSV, failedTickets []dto.InvalidTicketCSV) error {
	if err := writeProcessedToCSV(csv.validTicketsFilePath, processedTickets); err != nil {
		return err
	}
	if err := writeFailedToCSV(csv.invalidTicketsFilePath, failedTickets); err != nil {
		return err
	}
	return nil
}

func writeProcessedToCSV(fileName string, tickets []dto.ValidTicketCSV) error {
	file, err := createNewFile(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(&tickets, file)
}

func writeFailedToCSV(fileName string, tickets []dto.InvalidTicketCSV) error {
	file, err := createNewFile(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(&tickets, file)
}

func createNewFile(fileName string) (*os.File, error) {
	err := os.Remove(fileName)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
}
