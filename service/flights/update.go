package flights

import (
	"context"
	"sahaj/flight/dto"
	"sahaj/flight/entities"
	"sahaj/flight/reader"
	"sahaj/flight/writer"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type updater struct {
	reader    reader.Reader
	writer    writer.Writer
	validator CustomValidator
}

func NewUpadter(reader reader.Reader,
	writer writer.Writer,
	validator CustomValidator) Flights {
	return updater{
		reader:    reader,
		writer:    writer,
		validator: validator,
	}
}

func (u updater) Update(ctx context.Context) error {
	tickets, err := u.reader.Read()
	if err != nil {
		logrus.Error(err)
		return err
	}

	validTickets, invalidTickets := u.validateTickets(mapTicketCSVToTicketEntity(tickets), tickets)

	errs, _ := errgroup.WithContext(ctx)

	errs.Go(func() error {
		return u.writer.WriteValidTickets(validTickets)
	})

	errs.Go(func() error {
		return u.writer.WriteInvalidTickets(invalidTickets)
	})

	return errs.Wait()
}

func (u updater) validateTickets(transformedTickets []entities.Ticket, tickets []dto.TicketCSV) ([]dto.ValidTicketCSV, []dto.InvalidTicketCSV) {
	invalidTickets := []dto.InvalidTicketCSV{}
	validTickets := []dto.ValidTicketCSV{}
	for index, ticket := range transformedTickets {
		if err := u.validator.Validate(ticket); err != nil {
			logrus.WithField("first_name", tickets[index].FirstName).
				WithField("last_name", tickets[index].LastName).
				Info(err)

			invalidTickets = append(invalidTickets, createInvalidTicket(tickets[index], err.Error()))
			continue
		}
		validTickets = append(validTickets, createValidTicket(ticket))
	}
	return validTickets, invalidTickets
}

func createValidTicket(ticket entities.Ticket) dto.ValidTicketCSV {
	validTicket := mapTicketEntityToValidTicketCSV(ticket)
	validTicket.SetDiscountCode()
	return validTicket
}

func createInvalidTicket(ticket dto.TicketCSV, errMsg string) dto.InvalidTicketCSV {
	invalidTicket := mapTicketEntityToInvalidTicketCSV(ticket)
	invalidTicket.SetError(errMsg)
	return invalidTicket
}
