package heap

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	actual := []int{9, 4, 4, 1, 0, 53, 3, 102}

	Sort(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error")
	}
}
