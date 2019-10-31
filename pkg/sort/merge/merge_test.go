package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	actual := []int{4, 1, 3, 2}
	merge(actual, 0, 1, 2)

	if !reflect.DeepEqual(actual, []int{1, 4, 3, 2}) {
		t.Fatal("Error")
	}

	merge(actual, 2, 3, 4)

	if !reflect.DeepEqual(actual, []int{1, 4, 2, 3}) {
		t.Fatal("Error")
	}

	merge(actual, 0, 2, 4)

	if !reflect.DeepEqual(actual, []int{1, 2, 3, 4}) {
		t.Fatal("Error")
	}
}

func TestSort(t *testing.T) {
	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	actual := []int{9, 4, 4, 1, 0, 53, 3, 102}

	Sort(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error")
	}
}
