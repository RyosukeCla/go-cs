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
	input := []byte("lorem ipsum")

	encoder := NewEncoder(100)
	encoder.Write(input)
	codes := make([]Code, 1000)
	n, _ := encoder.Read(codes)
	codes = codes[:n]

	decoder := NewDecoder()
	decoder.Write(codes)
	output := make([]byte, 1000)
	n, _ = decoder.Read(output)
	output = output[:n]

	if !reflect.DeepEqual(input, output) {
		t.Fatal("Fail", input, output)
	}
}
