// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v2.go at 2019-08-15 19:24:55.9801922 +0300 EEST m=+0.008961201
package messages

import (
	"testing"
)

func TestReadFormat08V2Valid(t *testing.T) {

	data := buildValidBDS06V2Message()
	data[0] = data[0] | 0x40

	msg, err := ReadFormat08V2(false, false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS06V2Valid(t, msg)
}

func TestReadFormat08V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V2Message()[:6]
	data[0] = data[0] | 0x40

	_, err := ReadFormat08V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat08V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V2Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat08V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}
