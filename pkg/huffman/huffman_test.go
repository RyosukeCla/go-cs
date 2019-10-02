package huffman

import (
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	expected := []Code{
		{Symbol: "e", Code: []byte{0}},
		{Symbol: "a", Code: []byte{1, 0}},
		{Symbol: "d", Code: []byte{1, 1, 0}},
		{Symbol: "c", Code: []byte{1, 1, 1, 0}},
		{Symbol: "b", Code: []byte{1, 1, 1, 1, 0}},
		{Symbol: "f", Code: []byte{1, 1, 1, 1, 1}},
	}

	actual := Build([]Symbol{
		{Symbol: "a", Occurrence: 5},
		{Symbol: "b", Occurrence: 2},
		{Symbol: "c", Occurrence: 3},
		{Symbol: "d", Occurrence: 4},
		{Symbol: "e", Occurrence: 6},
		{Symbol: "f", Occurrence: 1},
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail", "expected:", expected, "actual:", actual)
	}
}
