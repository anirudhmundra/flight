package entities

import "time"

type Ticket struct {
	FirstName    string `validate:"alpha"`
	LastName     string `validate:"alpha"`
	Email        string `validate:"email"`
	MobileNumber string `validate:"numeric,len=10"`

	NumberOfPassengers uint
	PNR                string        `validate:"alphanum,len=6"`
	FareClass          string        `validate:"alpha,uppercase,len=1"`
	CabinCategory      CabinCategory `validate:"classCategory"`

	BookingDate time.Time `validate:"ltecsfield=TravelDate"`
	TravelDate  time.Time
}
