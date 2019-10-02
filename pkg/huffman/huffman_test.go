package huffman

import (
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	expected := []Code{
		{Symbol: byte(5), Code: []byte{0}},
		{Symbol: byte(1), Code: []byte{1, 0}},
		{Symbol: byte(4), Code: []byte{1, 1, 0}},
		{Symbol: byte(3), Code: []byte{1, 1, 1, 0}},
		{Symbol: byte(2), Code: []byte{1, 1, 1, 1, 0}},
		{Symbol: byte(6), Code: []byte{1, 1, 1, 1, 1}},
	}

	actual := Build([]Symbol{
		{Symbol: byte(1), Occurrence: 5},
		{Symbol: byte(2), Occurrence: 2},
		{Symbol: byte(3), Occurrence: 3},
		{Symbol: byte(4), Occurrence: 4},
		{Symbol: byte(5), Occurrence: 6},
		{Symbol: byte(6), Occurrence: 1},
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail", "expected:", expected, "actual:", actual)
	}
}
