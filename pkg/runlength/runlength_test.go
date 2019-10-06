package runlength

import (
	"reflect"
	"testing"
)

func TestEncoder(t *testing.T) {
	en := NewEncoder()
	en.Write([]byte("AAAAABBCCD"))
	codes := make([]Code, 8)
	n, _ := en.Read(codes)
	codes = codes[:n]

	expected := []Code{
		{Symbol: []byte("A")[0], Count: 5},
		{Symbol: []byte("B")[0], Count: 2},
		{Symbol: []byte("C")[0], Count: 2},
		{Symbol: []byte("D")[0], Count: 1},
	}
	if !reflect.DeepEqual(codes, expected) {
		t.Fatal("Error", codes)
	}
}

func TestDecoder(t *testing.T) {
	codes := []Code{
		{Symbol: []byte("A")[0], Count: 5},
		{Symbol: []byte("B")[0], Count: 2},
		{Symbol: []byte("C")[0], Count: 2},
		{Symbol: []byte("D")[0], Count: 1},
	}

	de := NewDecoder()
	de.Write(codes)

	actual := make([]byte, 100)
	n, _ := de.Read(actual)
	actual = actual[:n]

	expected := []byte("AAAAABBCCD")

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error")
	}
}
