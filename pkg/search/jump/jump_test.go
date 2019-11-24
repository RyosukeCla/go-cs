package jump

import (
	"testing"
)

func TestSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	if Search(array, 10) != 9 {
		t.Fatal("error")
	}
	if Search(array, 11) != 10 {
		t.Fatal("error")
	}
	if Search(array, 12) != -1 {
		t.Fatal("error")
	}
}
