package greedy

import "testing"

func Test(t *testing.T) {
	if Find(1, 2, 11) != 10 {
		t.Fatal("error")
	}

	if Find(5, 2, 1019) != 10 {
		t.Fatal("error")
	}
}
