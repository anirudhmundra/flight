package entities

import (
	"strings"
)

type CabinCategory int

const (
	Economy CabinCategory = iota + 1
	PremiumEconomy
	Business
	First
)

func (c CabinCategory) String() string {
	return [...]string{"Economy", "Premium Economy", "Business", "First"}[c-1]
}

func GetCabinCategory(category string) CabinCategory {
	category = strings.TrimSpace(category)
	switch category {
	case "Economy":
		return Economy
	case "Premium Economy":
		return PremiumEconomy
	case "Business":
		return Business
	case "First":
		return First
	default:
		return -1
	}
}
