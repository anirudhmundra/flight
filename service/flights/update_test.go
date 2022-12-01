package flights

import (
	"errors"
	"sahaj/flight/dto"
	"sahaj/flight/entities"
	"sahaj/flight/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/golang/mock/gomock"
)

type UpdateTestSuite struct {
	suite.Suite
	mockCtrl  *gomock.Controller
	reader    *mocks.MockReader
	writer    *mocks.MockWriter
	validator *mocks.MockCustomValidator
	updater   Flights
}

func (suite *UpdateTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.reader = mocks.NewMockReader(suite.mockCtrl)
	suite.writer = mocks.NewMockWriter(suite.mockCtrl)
	suite.validator = mocks.NewMockCustomValidator(suite.mockCtrl)
	suite.updater = NewUpadter(suite.reader, suite.writer, suite.validator)
}

func TestUpdateTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}

func (suite *UpdateTestSuite) TestUpdate_Successfully() {

	suite.reader.EXPECT().Read().Return([]dto.TicketCSV{
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
		},
	}, nil)

	suite.validator.EXPECT().Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
	}).Return(nil)

	suite.writer.EXPECT().Write([]dto.ValidTicketCSV{
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
	}, []dto.InvalidTicketCSV{}).Return(nil)

	err := suite.updater.Update()
	suite.Nil(err)
}

func (suite *UpdateTestSuite) TestUpdate_ReaderFails() {

	suite.reader.EXPECT().Read().Return([]dto.TicketCSV{}, errors.New("reader error"))

	err := suite.updater.Update()
	suite.Equal(errors.New("reader error"), err)
}

func (suite *UpdateTestSuite) TestUpdate_InvalidCabinCategory() {

	suite.reader.EXPECT().Read().Return([]dto.TicketCSV{
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
		},
	}, nil)

	suite.validator.EXPECT().Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      -1,

		BookingDate: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
	}).Return(errors.New("CabinCategory invalid"))

	suite.writer.EXPECT().Write([]dto.ValidTicketCSV{},
		[]dto.InvalidTicketCSV{
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
		}).Return(nil)

	err := suite.updater.Update()
	suite.Nil(err)
}

func (suite *UpdateTestSuite) TestUpdate_WriterFails() {

	suite.reader.EXPECT().Read().Return([]dto.TicketCSV{
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
		},
	}, nil)

	suite.validator.EXPECT().Validate(entities.Ticket{
		FirstName:    "firstname",
		LastName:     "lastname",
		Email:        "firstname@abc.com",
		MobileNumber: "9876543210",

		NumberOfPassengers: 12,
		PNR:                "1AB2C3",
		FareClass:          "A",
		CabinCategory:      entities.Economy,

		BookingDate: time.Date(2022, 1, 01, 0, 0, 0, 0, time.UTC),
		TravelDate:  time.Date(2021, 1, 01, 0, 0, 0, 0, time.UTC),
	}).Return(nil)

	suite.writer.EXPECT().Write([]dto.ValidTicketCSV{
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
	}, []dto.InvalidTicketCSV{}).Return(errors.New("error occurred"))

	err := suite.updater.Update()
	suite.Equal(errors.New("error occurred"), err)
}
