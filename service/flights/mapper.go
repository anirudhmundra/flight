package flights

import (
	"sahaj/flight/dto"
	"sahaj/flight/entities"
)

func mapTicketCSVToTicketEntity(tickets []dto.TicketCSV) []entities.Ticket {
	result := []entities.Ticket{}

	for _, ticket := range tickets {
		result = append(result, entities.Ticket{
			FirstName:          ticket.FirstName,
			LastName:           ticket.LastName,
			PNR:                ticket.PNR,
			FareClass:          ticket.FareClass,
			TravelDate:         ticket.TravelDate.Time,
			BookingDate:        ticket.BookingDate.Time,
			NumberOfPassengers: ticket.NumberOfPassengers,
			MobileNumber:       ticket.MobileNumber,
			Email:              ticket.Email,
			CabinCategory:      entities.GetCabinCategory(ticket.CabinCategory),
		})
	}
	return result
}

func mapTicketEntityToValidTicketCSV(ticket entities.Ticket, discountCode string) dto.ValidTicketCSV {
	return dto.ValidTicketCSV{
		FirstName:          ticket.FirstName,
		LastName:           ticket.LastName,
		PNR:                ticket.PNR,
		NumberOfPassengers: ticket.NumberOfPassengers,
		BookingDate: dto.DateTime{
			Time: ticket.BookingDate,
		},
		TravelDate: dto.DateTime{
			Time: ticket.TravelDate,
		},
		FareClass:     ticket.FareClass,
		Email:         ticket.Email,
		MobileNumber:  ticket.MobileNumber,
		CabinCategory: ticket.CabinCategory.String(),
		Discount:      discountCode,
	}
}

func mapTicketEntityToInvalidTicketCSV(ticket dto.TicketCSV, errMsg string) dto.InvalidTicketCSV {
	return dto.InvalidTicketCSV{
		FirstName:          ticket.FirstName,
		LastName:           ticket.LastName,
		PNR:                ticket.PNR,
		NumberOfPassengers: ticket.NumberOfPassengers,
		BookingDate:        ticket.BookingDate,
		TravelDate:         ticket.TravelDate,
		FareClass:          ticket.FareClass,
		Email:              ticket.Email,
		MobileNumber:       ticket.MobileNumber,
		CabinCategory:      ticket.CabinCategory,
		Error:              errMsg,
	}
}
