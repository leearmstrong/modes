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
	generateFile("format_09_v2_test.go", "Format09V2", "0x48", "false, ")
	generateFile("format_10_v2_test.go", "Format10V2", "0x50", "false, ")
	generateFile("format_11_v2_test.go", "Format11V2", "0x58", "false, ")
	generateFile("format_12_v2_test.go", "Format12V2", "0x60", "false, ")
	generateFile("format_13_v2_test.go", "Format13V2", "0x68", "false, ")
	generateFile("format_14_v2_test.go", "Format14V2", "0x70", "false, ")
	generateFile("format_15_v2_test.go", "Format15V2", "0x78", "false, ")
	generateFile("format_16_v2_test.go", "Format16V2", "0x80", "false, ")
	generateFile("format_17_v2_test.go", "Format17V2", "0x88", "false, ")
	generateFile("format_18_v2_test.go", "Format18V2", "0x90", "false, ")
	generateFile("format_20_v2_test.go", "Format20V2", "0xA0", "")
	generateFile("format_21_v2_test.go", "Format21V2", "0xA8", "")
	generateFile("format_22_v2_test.go", "Format22V2", "0xB0", "")
}

func generateFile(fileName string, name string, messageCode string, nicSupplementA string) {
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
			Timestamp      time.Time
			Name           string
			MessageCode    string
			NICSupplementA string
		}{
			Timestamp:      time.Now(),
			Name:           name,
			MessageCode:    messageCode,
			NICSupplementA: nicSupplementA,
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
// This file was generated by gen_tests_v2.go at {{ .Timestamp }}
package messages

import (
	"testing"
)

func TestRead{{ .Name }}Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | {{ .MessageCode }}

	msg, err := Read{{ .Name }}({{.NICSupplementA}}data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V2Valid(t, msg)
}

func TestRead{{ .Name }}TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V2Message()[:6]
	data[0] = data[0] | {{ .MessageCode }}

	_, err := Read{{ .Name }}({{.NICSupplementA}}data)
	if err == nil {
		t.Error(err)
	}
}

func TestRead{{ .Name }}BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x01

	_, err := Read{{ .Name }}({{.NICSupplementA}}data)
	if err == nil {
		t.Error(err)
	}
}
`))
