package reader

import (
	"sahaj/flight/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCSVRead_InvalidFilePath(t *testing.T) {
	reader := NewCSVReader("invalid")
	tickets, err := reader.Read()
	assert.Nil(t, tickets)
	assert.NotNil(t, err)
}

func TestCSVRead_Successfully(t *testing.T) {
	reader := NewCSVReader("testdata/tickets.csv")
	tickets, err := reader.Read()
	assert.Equal(t, []dto.TicketCSV{
		{
			FirstName: "Abhishek",
			LastName:  "Kumar",
			PNR:       "ABC123",
			FareClass: "F",
			TravelDate: dto.DateTime{
				Time: time.Date(2019, time.July, 31, 0, 0, 0, 0, time.UTC),
			},
			NumberOfPassengers: 2,
			BookingDate: dto.DateTime{
				Time: time.Date(2019, time.May, 21, 0, 0, 0, 0, time.UTC),
			},
			Email:         "abhishek@zzz.com",
			MobileNumber:  "9876543210",
			CabinCategory: "Economy"},
	}, tickets)
	assert.Nil(t, err)
}


func TestCSVRead_InvalidTickets(t *testing.T) {
	reader := NewCSVReader("testdata/invalid_tickets.csv")
	tickets, err := reader.Read()
	assert.Nil(t, tickets)
	assert.NotNil(t, err)
}
