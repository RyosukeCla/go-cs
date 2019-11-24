package binary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if Search(array, 3) != 2 {
		t.Fatal("error")
	}
}
