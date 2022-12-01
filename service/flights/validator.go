package flights

import (
	"errors"
	"fmt"
	"sahaj/flight/entities"
	"strings"

	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validate *validator.Validate
}

type CustomValidator interface {
	Validate(ticket entities.Ticket) error
}

func NewCustomValidtor() customValidator {
	validate := validator.New()
	validate.RegisterValidation("classCategory", func(fl validator.FieldLevel) bool {
		switch fl.Field().Interface().(entities.CabinCategory) {
		case entities.Economy, entities.PremiumEconomy, entities.Business, entities.First:
			return true
		default:
			return false
		}
	})
	return customValidator{
		validate: validate,
	}
}

func (v customValidator) Validate(ticket entities.Ticket) error {
	err := v.validate.Struct(&ticket)
	if err != nil {
		return generateCustomError(err)
	}
	return nil
}

func generateCustomError(err error) error {
	errMsgs := []string{}
	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		errMsgs = append(errMsgs, fmt.Sprintf("%s invalid", validationErr.Field()))
	}
	return errors.New(strings.Join(errMsgs, ","))
}
