// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2019-08-15 10:36:54.6288357 +0300 EEST m=+0.007968401
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat08V0Valid(t *testing.T) {

	data := buildValidBDS06V0Message()
	data[0] = data[0] | 0x40

	msg, err := ReadFormat08V0(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format08V0 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format08V0.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06V0Valid(t, msg)
}

func TestReadFormat08V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V0Message()[:6]
	data[0] = data[0] | 0x40

	_, err := ReadFormat08V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat08V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V0Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat08V0(data)
	if err == nil {
		t.Error(err)
	}
}
