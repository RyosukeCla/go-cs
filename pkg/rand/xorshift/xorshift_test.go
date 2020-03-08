package xorshift

import (
	"testing"
)

func Test(t *testing.T) {
	rand := New(123)
	for i := 0; i < 1000; i++ {
		r := rand.Generate()
		if 0 < r || r >= 1 {
			t.Fatal("error")
		}
	}
}
