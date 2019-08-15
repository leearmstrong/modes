// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates list_converter.go. It can be invoked by running
// go generate

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func main() {
	generateFile("format_05_v2.go", "Format05V2", "Format05")
	generateFile("format_06_v2.go", "Format06V2", "Format06")
	generateFile("format_07_v2.go", "Format07V2", "Format07")
	generateFile("format_08_v2.go", "Format08V2", "Format08")
}

func generateFile(fileName string, name string, formatName string) {
	// Open the target file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Close at the end
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatal(err)
		}
	}()

	// Execute the template
	err = builderTemplate.Execute(
		f,
		struct {
			Timestamp  time.Time
			Name       string
			FormatName string
		}{
			Timestamp:  time.Now(),
			Name:       name,
			FormatName: formatName,
		})
	if err != nil {
		log.Fatal(err)
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v2.go at {{ .Timestamp }}
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// ------------------------------------------------------------------------------------
//
//                                {{ .Name }}
//
// ------------------------------------------------------------------------------------

// {{ .Name }} is a message at the format BDS 0,6
type {{ .Name }} struct {
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
func (message {{ .Name }}) GetMessageFormat() adsb.MessageFormat {
	return adsb.{{ .FormatName }}
}

// GetRegister returns the register of the message
func (message {{ .Name }}) GetRegister() bds.Register {
	return adsb.{{ .FormatName }}.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message {{ .Name }}) GetSubtype() adsb.Subtype{
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message {{ .Name }}) GetMinimumADSBLevel() adsb.MessageLevel{
	return adsb.MessageLevel2
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message {{ .Name }}) GetMaximumADSBLevel() adsb.MessageLevel{
	return adsb.MessageLevel2
}

// GetMovement returns the Movement
func (message {{ .Name }}) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message {{ .Name }}) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message {{ .Name }}) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message {{ .Name }}) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message {{ .Name }}) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message {{ .Name }}) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message {{ .Name }}) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadiusV2
func (message {{ .Name }}) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadiusV2 {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message {{ .Name }}) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// ToString returns a basic, but readable, representation of the message
func (message {{ .Name }}) ToString() string {
	return messageBDS06V2ToString(message)
}

// Read{{ .Name }} reads a message at the format {{ .Name }} for ADSB V2
func Read{{ .Name }}(nicSupplementA bool, nicSupplementC bool, data []byte) (*{{ .Name }}, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.{{ .FormatName }}.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format {{ .Name }}", formatTypeCode)
	}

	hcr, nic := getHCRAndNICForV2(formatTypeCode, nicSupplementA, nicSupplementC)

	return &{{ .Name }}{
		Movement:                      fields.ReadMovement(data),
		GroundTrackStatus:             fields.ReadGroundTrackStatus(data),
		GroundTrack:                   fields.ReadGroundTrack(data),
		Time:                          fields.ReadTime(data),
		CPRFormat:                     fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:               fields.ReadEncodedLatitude(data),
		EncodedLongitude:              fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius:   hcr,
		NavigationIntegrityCategory:   nic,
	}, nil
}
`))
