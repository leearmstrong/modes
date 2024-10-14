package fields

import (
	"errors"
	"fmt"
)

// AircraftCategorySetD defines the type of the AircraftCategory
//
// Specified in Doc 9871 / Table A-2-8
type AircraftCategorySetD byte

const (
	// ACSDReserved0 indicates Reserved
	ACSDReserved0 AircraftCategorySetD = 0
	// ACSDReserved1 indicates Reserved
	ACSDReserved1 AircraftCategorySetD = 1
	// ACSDReserved2 indicates Reserved
	ACSDReserved2 AircraftCategorySetD = 2
	// ACSDReserved3 indicates Reserved
	ACSDReserved3 AircraftCategorySetD = 3
	// ACSDReserved4 indicates Reserved
	ACSDReserved4 AircraftCategorySetD = 4
	// ACSDReserved5 indicates Reserved
	ACSDReserved5 AircraftCategorySetD = 5
	// ACSDReserved6 indicates Reserved
	ACSDReserved6 AircraftCategorySetD = 6
	// ACSDReserved7 indicates Reserved
	ACSDReserved7 AircraftCategorySetD = 7
)

// GetCategorySetName returns the name of the category set
func (category AircraftCategorySetD) GetCategorySetName() string {
	return "Set D (Reserved)"
}

// ToString returns a basic, but readable, representation of the field
func (category AircraftCategorySetD) ToString() string {

	switch category {
	case ACSDReserved0:
		return "0 - Reserved"
	case ACSDReserved1:
		return "1 - Reserved"
	case ACSDReserved2:
		return "2 - Reserved"
	case ACSDReserved3:
		return "3 - Reserved"
	case ACSDReserved4:
		return "4 - Reserved"
	case ACSDReserved5:
		return "5 - Reserved"
	case ACSDReserved6:
		return "6 - Reserved"
	case ACSDReserved7:
		return "7 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ToShortString returns a basic category and type concatenated
func (category AircraftCategorySetD) ToShortString() string {

	switch category {
	case ACSDReserved0:
		return "D0"
	case ACSDReserved1:
		return "D1"
	case ACSDReserved2:
		return "D2"
	case ACSDReserved3:
		return "D3"
	case ACSDReserved4:
		return "D4"
	case ACSDReserved5:
		return "D5"
	case ACSDReserved6:
		return "D6"
	case ACSDReserved7:
		return "D7"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// CheckCoherency checks that the Category is coherent
func (category AircraftCategorySetD) CheckCoherency() error {

	return errors.New("field FormatTypeCode is 1 (Category set D) which is a reserved Category set")
}
