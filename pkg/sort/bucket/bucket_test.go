package bucket

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	array := []int{9, 4, 4, 1, 0, 53, 3, 102}

	actual := Sort(array, 103)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error")
	}
}
