package midsquare

import (
	"testing"
)

func Test(t *testing.T) {
	rand := New(4, 9090)
	for i := 0; i < 1000; i++ {
		r := rand.Generate()
		if 0 < r || r >= 1 {
			t.Fatal("error")
		}
	}
}
