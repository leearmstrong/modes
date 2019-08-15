package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19GroundNormal is a message at the format BDS 9,0
type Format19GroundNormal struct {
	IntentChange                  fields.IntentChange
	IFRCapability                 fields.IFRCapability
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	DirectionEastWest             fields.DirectionEastWest
	VelocityEWNormal              fields.VelocityEWNormal
	DirectionNorthSouth           fields.DirectionNorthSouth
	VelocityNSNormal              fields.VelocityNSNormal
	VerticalRateSource            fields.VerticalRateSource
	VerticalRateSign              fields.VerticalRateSign
	VerticalRate                  fields.VerticalRate
	DifferenceGNSSBaroSign        fields.DifferenceGNSSBaroSign
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format19GroundNormal) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format19V0OrMore
}

// GetRegister returns the register of the message
func (message *Format19GroundNormal) GetRegister() bds.Register {
	return adsb.Format19V0OrMore.GetRegister()
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message *Format19GroundNormal) GetSubtype() fields.Subtype {
	return fields.SubtypeGroundSpeedNormal
}

// GetIntentChange returns the IntentChange
func (message *Format19GroundNormal) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message *Format19GroundNormal) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message *Format19GroundNormal) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message *Format19GroundNormal) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message *Format19GroundNormal) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message *Format19GroundNormal) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message *Format19GroundNormal) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message *Format19GroundNormal) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19GroundNormal) ToString() string {
	return fmt.Sprintf("Message:                         %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change:                   %v\n"+
		"IFR Capability:                  %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Direction EW:                    %v\n"+
		"Velocity EW Normal:              %v\n"+
		"Direction NS:                    %v\n"+
		"Velocity NS Normal:              %v\n"+
		"Vertical Rate Source:            %v\n"+
		"Vertical Rate Sign:              %v\n"+
		"Vertical Rate:                   %v\n"+
		"Difference GNSS Baro Sign:       %v\n"+
		"Difference GNSS Baro:            %v",
		adsb.Format19V0OrMore.ToString(),
		message.GetSubtype().ToString(),
		message.IntentChange.ToString(),
		message.IFRCapability.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.DirectionEastWest.ToString(),
		message.VelocityEWNormal.ToString(),
		message.DirectionNorthSouth.ToString(),
		message.VelocityNSNormal.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadFormat19GroundNormal reads a message at the format Format19 / Subtype 1 (Ground speed normal)
func ReadFormat19GroundNormal(data []byte) (*Format19GroundNormal, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format19V0OrMore.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat19GroundNormal", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedNormal {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat19GroundNormal", subType.ToString())
	}

	return &Format19GroundNormal{
		IntentChange:                  fields.ReadIntentChange(data),
		IFRCapability:                 fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		DirectionEastWest:             fields.ReadDirectionEastWest(data),
		VelocityEWNormal:              fields.ReadVelocityEWNormal(data),
		DirectionNorthSouth:           fields.ReadDirectionNorthSouth(data),
		VelocityNSNormal:              fields.ReadVelocityNSNormal(data),
		VerticalRateSource:            fields.ReadVerticalRateSource(data),
		VerticalRateSign:              fields.ReadVerticalRateSign(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		DifferenceGNSSBaroSign:        fields.ReadDifferenceGNSSBaroSign(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
