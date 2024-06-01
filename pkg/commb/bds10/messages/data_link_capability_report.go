package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/commb"
	"github.com/twuillemin/modes/pkg/commb/bds10/fields"
)

// DataLinkCapabilityReport is a message at the format BDS 1,0
//
// Specified in Doc 9871 / D.2.4.1
type DataLinkCapabilityReport struct {
	ContinuationFlag                     fields.ContinuationFlag
	OverlayCommandCapability             fields.OverlayCommandCapability
	ACASStatus                           fields.ACASStatus
	ModeSSubnetworkVersion               fields.ModeSSubnetworkVersion
	TransponderEnhancedProtocolIndicator fields.TransponderEnhancedProtocolIndicator
	ModeSSpecificServicesCapability      fields.ModeSSpecificServicesCapability
	UplinkELMThroughputCapability        fields.UplinkELMThroughputCapability
	DownlinkELMThroughputCapability      fields.DownlinkELMThroughputCapability
	AircraftIdentificationCapability     fields.AircraftIdentificationCapability
	SquitterCapabilitySubfield           fields.SquitterCapabilitySubfield
	SurveillanceIdentifierCode           fields.SurveillanceIdentifierCode
	CommonUsageGICBCapability            fields.CommonUsageGICBCapability
	ACASHybridSurveillanceCapability     fields.ACASHybridSurveillanceCapability
	ACASGenerationCapability             fields.ACASGenerationCapability
	ACASApplicableDocument               fields.ACASApplicableDocument
	DTESubAddressStatuses                fields.DTESubAddressStatuses
}

// GetMessageFormat returns the register of the message
func (message DataLinkCapabilityReport) GetMessageFormat() commb.MessageFormat {
	return commb.FormatDataLinkCapabilityReport
}

// ToString returns a basic, but readable, representation of the message
func (message DataLinkCapabilityReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Continuation Flag:                       %v\n"+
		"ACAS Status                              %v\n"+
		"ModeS Subnetwork Version                 %v\n"+
		"Transponder Enhanced Protocol Indicator  %v\n"+
		"ModeS Specific Services Capability       %v\n"+
		"Uplink ELM Throughput Capability         %v\n"+
		"Downlink ELM Throughput Capability       %v\n"+
		"Aircraft Identification Capability       %v\n"+
		"Squitter Capability Subfield             %v\n"+
		"Surveillance Identifier Code             %v\n"+
		"Common Usage GICB Capability             %v\n"+
		"ACAS Hybrid Surveillance Capability      %v\n"+
		"ACAS Generation Capability               %v\n"+
		"ACAS Applicable Document                 %v\n"+
		"DTE SubAddress Statuses                  %v\n",
		commb.GetMessageFormatInformation(&message),
		message.OverlayCommandCapability.ToString(),
		message.ACASStatus.ToString(),
		message.ModeSSubnetworkVersion.ToString(),
		message.TransponderEnhancedProtocolIndicator.ToString(),
		message.ModeSSpecificServicesCapability.ToString(),
		message.UplinkELMThroughputCapability.ToString(),
		message.DownlinkELMThroughputCapability.ToString(),
		message.AircraftIdentificationCapability.ToString(),
		message.SquitterCapabilitySubfield.ToString(),
		message.SurveillanceIdentifierCode.ToString(),
		message.CommonUsageGICBCapability.ToString(),
		message.ACASHybridSurveillanceCapability.ToString(),
		message.ACASGenerationCapability.ToString(),
		message.ACASApplicableDocument.ToString(),
		message.DTESubAddressStatuses.ToString())
}

// ReadDataLinkCapabilityReport reads a message as a DataLinkCapabilityReport
func ReadDataLinkCapabilityReport(data []byte) (*DataLinkCapabilityReport, error) {
	err := CheckIfDataReadable(data)
	if err != nil {
		return nil, err
	}

	return &DataLinkCapabilityReport{
		ContinuationFlag:                     fields.ReadContinuationFlag(data),
		OverlayCommandCapability:             fields.ReadOverlayCommandCapability(data),
		ACASStatus:                           fields.ReadACASStatus(data),
		ModeSSubnetworkVersion:               fields.ReadModeSSubnetworkVersion(data),
		TransponderEnhancedProtocolIndicator: fields.ReadTransponderEnhancedProtocolIndicator(data),
		ModeSSpecificServicesCapability:      fields.ReadModeSSpecificServicesCapability(data),
		UplinkELMThroughputCapability:        fields.ReadUplinkELMThroughputCapability(data),
		DownlinkELMThroughputCapability:      fields.ReadDownlinkELMThroughputCapability(data),
		AircraftIdentificationCapability:     fields.ReadAircraftIdentificationCapability(data),
		SquitterCapabilitySubfield:           fields.ReadSquitterCapabilitySubfield(data),
		SurveillanceIdentifierCode:           fields.ReadSurveillanceIdentifierCode(data),
		CommonUsageGICBCapability:            fields.ReadCommonUsageGICBCapability(data),
		ACASHybridSurveillanceCapability:     fields.ReadACASHybridSurveillanceCapability(data),
		ACASGenerationCapability:             fields.ReadACASGenerationCapability(data),
		ACASApplicableDocument:               fields.ReadACASApplicableDocument(data),
		DTESubAddressStatuses:                fields.ReadDTESubAddressStatuses(data),
	}, nil
}

// CheckIfDataReadable checks if the data can be read as a DataLinkCapabilityReport
func CheckIfDataReadable(data []byte) error {
	if len(data) != 7 {
		return errors.New("the data for Comm-B DataLinkCapabilityReport message must be 7 bytes long")
	}

	// First byte is simply the BDS format 0001 0000
	if data[0] != 0x10 {
		return errors.New("the first byte of data is not 0x10")
	}

	// Bits 10 to 14 are reserved and must be 0
	if data[1]&0x7C != 0 {
		return errors.New("the bits 10 to 14 are reserved and must be 0")
	}

	return nil
}
