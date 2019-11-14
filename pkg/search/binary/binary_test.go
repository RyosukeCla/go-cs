package binary

import "testing"

func TestSearch(t *testing.T) {
	array := []int{100, 13, 34, 1, 5, 999, 65}
	if Search(array, 1) != 3 {
		t.Fatal("error")
	}
}
