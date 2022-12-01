package flights

import (
	"sahaj/flight/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Successful(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Nil(t, err)
}

func TestValidator_InvalidFirstName(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firs12tname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "FirstName invalid", err.Error())
}

func TestValidator_EmptyLastName(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "LastName invalid", err.Error())
}

func TestValidator_InvalidEmail(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstnameabc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "Email invalid", err.Error())
}

func TestValidator_InvalidMobileNumber(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "98765a43210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "MobileNumber invalid", err.Error())
}

func TestValidator_InvalidPNR(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1A2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "PNR invalid", err.Error())
}

func TestValidator_InvalidFareClass(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "a",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "FareClass invalid", err.Error())
}

func TestValidator_InvalidCabinCategory(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      100,

		BookingDate: time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "CabinCategory invalid", err.Error())
}

func TestValidator_InvalidBookingDate(t *testing.T) {
	customValidator := NewCustomValidtor()

	err := customValidator.Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Business,

		BookingDate: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
	})
	assert.Equal(t, "BookingDate invalid", err.Error())
}
