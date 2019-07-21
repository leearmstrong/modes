package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS08 is the basic interface that ADSB messages at the format BDS 0,8 are expected to implement
type MessageBDS08 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetAircraftCategory returns the aircraft category
	GetAircraftCategory() fields.AircraftCategory
	// GetAircraftIdentification returns the identity of the aircraft
	GetAircraftIdentification() fields.AircraftIdentification
}

var bds08Code = "BDS 0,8"
var bds08Name = "Extended squitter aircraft identification and category"

func bds08ToString(message MessageBDS08) string {
	return fmt.Sprintf("Message:                 %v - %v (%v)\n"+
		"Aircraft Category:       %v (%v)\n"+
		"Aircraft Identification: %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetAircraftCategory().ToString(),
		message.GetAircraftCategory().GetCategorySetName(),
		message.GetAircraftIdentification())
}

// ReadBDS08 reads a message at the format BDS 0,8. As this format does not have changes from ADSB V0 to
// ADSB V2, the returned ADSBLevel is always the given one.
//
// Params:
//    - adsbLevel: The ADSB level request (not used, but present for coherency)
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS08(adsbLevel common.ADSBLevel, data []byte) (MessageBDS08, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1:
		message, err := readFormat01(data)
		return message, adsbLevel, err
	case 2:
		message, err := readFormat02(data)
		return message, adsbLevel, err
	case 3:
		message, err := readFormat03(data)
		return message, adsbLevel, err
	case 4:
		message, err := readFormat04(data)
		return message, adsbLevel, err
	}

	return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,8 format", formatTypeCode)
}
