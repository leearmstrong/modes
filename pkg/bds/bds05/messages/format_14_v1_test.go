// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2024-06-01 14:32:25.0614957 +0300 EEST m=+0.003706101
package messages

import (
	"testing"
)

func TestReadFormat14V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x70

	msg, err := ReadFormat14V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V1Valid(t, msg)
}

func TestReadFormat14V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V1Message()[:6]
	data[0] = data[0] | 0x70

	_, err := ReadFormat14V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat14V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat14V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
