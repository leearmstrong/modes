// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2024-06-01 12:49:20.2240277 +0300 EEST m=+0.002700901
package messages

import (
	"testing"
)

func TestReadFormat07V1Valid(t *testing.T) {

	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x38

	msg, err := ReadFormat07V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS06V1Valid(t, msg)
}

func TestReadFormat07V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V1Message()[:6]
	data[0] = data[0] | 0x38

	_, err := ReadFormat07V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat07V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat07V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
