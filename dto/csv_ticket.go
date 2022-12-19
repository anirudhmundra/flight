package dto

import (
	"time"
)

type TicketCSV struct {
	FirstName          string   `csv:"First_name"`
	LastName           string   `csv:"Last_name"`
	PNR                string   `csv:"PNR"`
	FareClass          string   `csv:"Fare_class"`
	TravelDate         DateTime `csv:"Travel_date"`
	NumberOfPassengers uint     `csv:"Pax"`
	BookingDate        DateTime `csv:"Ticketing_date"`
	Email              string   `csv:"Email"`
	MobileNumber       string   `csv:"Mobile_phone"`
	CabinCategory      string   `csv:"Booked_cabin"`
}

type ValidTicketCSV struct {
	FirstName          string   `csv:"First_name"`
	LastName           string   `csv:"Last_name"`
	PNR                string   `csv:"PNR"`
	FareClass          string   `csv:"Fare_class"`
	TravelDate         DateTime `csv:"Travel_date"`
	NumberOfPassengers uint     `csv:"Pax"`
	BookingDate        DateTime `csv:"Ticketing_date"`
	Email              string   `csv:"Email"`
	MobileNumber       string   `csv:"Mobile_phone"`
	CabinCategory      string   `csv:"Booked_cabin"`
	Discount           string   `csv:"Discount_code"`
}

func (t *ValidTicketCSV) SetDiscountCode() {
	t.Discount = getDiscountCodes.list[t.FareClass]
}

type InvalidTicketCSV struct {
	FirstName          string   `csv:"First_name"`
	LastName           string   `csv:"Last_name"`
	PNR                string   `csv:"PNR"`
	FareClass          string   `csv:"Fare_class"`
	TravelDate         DateTime `csv:"Travel_date"`
	NumberOfPassengers uint     `csv:"Pax"`
	BookingDate        DateTime `csv:"Ticketing_date"`
	Email              string   `csv:"Email"`
	MobileNumber       string   `csv:"Mobile_phone"`
	CabinCategory      string   `csv:"Booked_cabin"`
	Error              string   `csv:"Error"`
}

func (t *InvalidTicketCSV) SetError(msg string) {
	t.Error = msg
}

type DateTime struct {
	time.Time
}

func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

func (date *DateTime) UnmarshalCSV(csvDate string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csvDate)
	return err
}
