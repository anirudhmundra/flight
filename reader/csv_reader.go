package reader

import (
	"os"
	"sahaj/flight/dto"

	"github.com/gocarina/gocsv"
)

type csvReader struct {
	fileName string
}

func NewCSVReader(fileName string) Reader {
	return csvReader{fileName: fileName}
}

func (csv csvReader) Read() ([]dto.TicketCSV, error) {
	file, err := os.OpenFile(csv.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tickets := []dto.TicketCSV{}

	if err := gocsv.UnmarshalFile(file, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}
