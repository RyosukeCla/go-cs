package lz77

import (
	"reflect"
	"testing"
)

var windowSize int = 4
var expectedInput []byte = []byte{0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2}
var expectedCodes []Code = []Code{
	{Dist: 0, Len: 0, Token: []byte{0}},
	{Dist: 0, Len: 0, Token: []byte{1}},
	{Dist: 0, Len: 0, Token: []byte{2}},
	{Dist: 3, Len: 3, Token: []byte{}},
	{Dist: 3, Len: 3, Token: []byte{}},
	{Dist: 3, Len: 3, Token: []byte{}},
}

func TestEncoder(t *testing.T) {
	encoder := NewEncoder(windowSize)
	encoder.Write(expectedInput)
	codes := make([]Code, 20)
	n, _ := encoder.Read(codes)
	actualCodes := codes[:n]

	if !reflect.DeepEqual(actualCodes, expectedCodes) {
		t.Fatal("Fail", actualCodes, expectedCodes)
	}
}

func TestDecoder(t *testing.T) {
	decoder := NewDecoder()
	decoder.Write(expectedCodes)
	decoded := make([]byte, 20)
	n, _ := decoder.Read(decoded)
	actualInput := decoded[:n]

	if !reflect.DeepEqual(actualInput, expectedInput) {
		t.Fatal("Fail", actualInput, expectedInput)
	}
}

func TestEncoderAndDecoder(t *testing.T) {
	input := []byte(`
		Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
		Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
		Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	`)

	encoder := NewEncoder(100)
	encoder.Write(input)
	codes := make([]Code, 500)
	n, _ := encoder.Read(codes)
	codes = codes[:n]

	decoder := NewDecoder()
	decoder.Write(codes)
	output := make([]byte, 500)
	n, _ = decoder.Read(output)
	output = output[:n]

	if !reflect.DeepEqual(input, output) {
		t.Fatal("Fail", input, output)
	}
}
