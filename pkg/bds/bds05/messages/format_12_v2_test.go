// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v2.go at 2024-06-01 14:32:25.6396976 +0300 EEST m=+0.003348501
package messages

import (
	"testing"
)

func TestReadFormat12V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x60

	msg, err := ReadFormat12V2(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V2Valid(t, msg)
}

func TestReadFormat12V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V2Message()[:6]
	data[0] = data[0] | 0x60

	_, err := ReadFormat12V2(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat12V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat12V2(false, data)
	if err == nil {
		t.Error(err)
	}
}
