package writer

import (
	"sahaj/flight/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCSVWriterValidTickets_Successfully(t *testing.T) {
	writer := NewCSVWriter("testdata/valid.csv", "testdata/invalid.csv")
	err := writer.WriteValidTickets([]dto.ValidTicketCSV{
		{
			FirstName:    "firstname",
			LastName:     "lastname",
			Email:        "firstname@abc.com",
			MobileNumber: "9876543210",

			NumberOfPassengers: 12,
			PNR:                "1AB2C3",
			FareClass:          "A",
			CabinCategory:      "Economy",

			BookingDate: dto.DateTime{
				Time: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
			},
			TravelDate: dto.DateTime{
				Time: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
			},
			Discount: "OFFER_20",
		},
	})
	assert.Nil(t, err)
}

func TestCSVWriterInvalidTickets_Successfully(t *testing.T) {
	writer := NewCSVWriter("testdata/valid.csv", "testdata/invalid.csv")
	err := writer.WriteInvalidTickets([]dto.InvalidTicketCSV{
		{
			FirstName:    "firstname",
			LastName:     "lastname",
			Email:        "firstname@abc.com",
			MobileNumber: "9876543210",

			NumberOfPassengers: 12,
			PNR:                "1AB2C3",
			FareClass:          "A",
			CabinCategory:      "Economy1",

			BookingDate: dto.DateTime{
				Time: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
			},
			TravelDate: dto.DateTime{
				Time: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
			},
			Error: "CabinCategory invalid",
		},
	})
	assert.Nil(t, err)
}
