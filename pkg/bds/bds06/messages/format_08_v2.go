// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v2.go at 2024-06-01 14:32:27.3752567 +0300 EEST m=+0.003534101
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format08V2
//
// ------------------------------------------------------------------------------------

// Format08V2 is a message at the format BDS 0,6
type Format08V2 struct {
	Movement                    fields.Movement
	GroundTrackStatus           fields.GroundTrackStatus
	GroundTrack                 fields.GroundTrack
	Time                        fields.Time
	CPRFormat                   fields.CompactPositionReportingFormat
	EncodedLatitude             fields.EncodedLatitude
	EncodedLongitude            fields.EncodedLongitude
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusV2
	NavigationIntegrityCategory byte
}

// GetMessageFormat returns the ADSB format of the message
func (message Format08V2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format08
}

// GetRegister returns the register of the message
func (message Format08V2) GetRegister() bds.Register {
	return adsb.Format08.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format08V2) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message Format08V2) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message Format08V2) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetMovement returns the Movement
func (message Format08V2) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message Format08V2) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message Format08V2) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message Format08V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message Format08V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message Format08V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message Format08V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadiusV2
func (message Format08V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadiusV2 {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message Format08V2) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// ToString returns a basic, but readable, representation of the message
func (message Format08V2) ToString() string {
	return messageBDS06V2ToString(message)
}

// ReadFormat08V2 reads a message at the format Format08V2 for ADSB V2
func ReadFormat08V2(nicSupplementA bool, nicSupplementC bool, data []byte) (*Format08V2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format08.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format08V2", formatTypeCode)
	}

	hcr, nic := getHCRAndNICForV2(formatTypeCode, nicSupplementA, nicSupplementC)

	return &Format08V2{
		Movement:                    fields.ReadMovement(data),
		GroundTrackStatus:           fields.ReadGroundTrackStatus(data),
		GroundTrack:                 fields.ReadGroundTrack(data),
		Time:                        fields.ReadTime(data),
		CPRFormat:                   fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:             fields.ReadEncodedLatitude(data),
		EncodedLongitude:            fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}
