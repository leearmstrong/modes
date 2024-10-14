package fields

import (
	"fmt"
)

// Subtype is the Airborne Velocity Subtype Code definition

type AirSpeedType byte

const (
	// AirSpeedTypeIAS indicates IAS
	AirSpeedTypeIAS AirSpeedType = 0
	// AirSpeedTypeTAS indicates TAS
	AirSpeedTypeTAS AirSpeedType = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit AirSpeedType) ToString() string {

	switch bit {
	case AirSpeedTypeIAS:
		return "0 - IAS"
	case AirSpeedTypeTAS:
		return "1 - TAS"
	default:
		return fmt.Sprintf("%v - Unknown type", bit)
	}
}

// ReadAirspeedType reads the Airspeed Type from a single bit
func ReadAirspeedType(data []byte) AirSpeedType {
	bits := data[3] & 0x80 >> 7
	return AirSpeedType(bits)
}
