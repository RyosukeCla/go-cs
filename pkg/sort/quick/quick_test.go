package quick

import (
	"reflect"
	"testing"
)

func TestPartition_1(t *testing.T) {
	array := []int{3, 2, 1, 5, 4, 6}
	actual := partition(array, 0, len(array)-1)
	expected := 2
	if actual != expected {
		t.Fatalf("Error: actual %d, expected %d", actual, expected)
	}
	expectedArray := []int{1, 2, 3, 5, 4, 6}
	if !reflect.DeepEqual(array, expectedArray) {
		t.Fatalf("Error: actual %d, expected %d", array, expectedArray)
	}
}

func TestPartition_2(t *testing.T) {
	array := []int{2, 4, 1, 5}
	actual := partition(array, 0, len(array)-1)
	expected := 1
	if actual != expected {
		t.Fatalf("Error: actual %d, expected %d", actual, expected)
	}
	expectedArray := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(array, expectedArray) {
		t.Fatalf("Error: actual %d, expected %d", array, expectedArray)
	}
}

func TestSort(t *testing.T) {
	actual := []int{9, 4, 4, 1, 0, 53, 3, 102}

	Sort(actual)

	expected := []int{0, 1, 3, 4, 4, 9, 53, 102}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Error", actual)
	}
}
