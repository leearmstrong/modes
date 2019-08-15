// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2019-08-15 10:36:55.1465472 +0300 EEST m=+0.008955401
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat06V1Valid(t *testing.T) {

	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x30

	msg, err := ReadFormat06V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format06V1 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format06V1.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06V1Valid(t, msg)
}

func TestReadFormat06V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V1Message()[:6]
	data[0] = data[0] | 0x30

	_, err := ReadFormat06V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat06V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat06V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
