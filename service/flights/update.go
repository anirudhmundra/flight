package flights

import (
	"sahaj/flight/dto"
	"sahaj/flight/reader"
	"sahaj/flight/writer"

	"github.com/sirupsen/logrus"
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

func (u updater) Update() error {
	tickets, err := u.reader.Read()
	if err != nil {
		logrus.Error(err)
		return err
	}
	transformedTickets := mapTicketCSVToTicketEntity(tickets)

	failedTickets := []dto.InvalidTicketCSV{}
	processedTickets := []dto.ValidTicketCSV{}
	for index, ticket := range transformedTickets {
		if err := u.validator.Validate(ticket); err != nil {
			logrus.WithField("first_name", tickets[index].FirstName).
				WithField("last_name", tickets[index].LastName).
				Info(err)

			failedTickets = append(failedTickets, mapTicketEntityToFailedTicketCSV(tickets[index], err.Error()))
			continue
		}
		processedTickets = append(processedTickets, mapTicketEntityToProcessedTicketCSV(ticket, getDiscountCodes.list[ticket.FareClass]))
	}

	if err := u.writer.Write(processedTickets, failedTickets); err != nil {
		return err
	}
	return nil
}
