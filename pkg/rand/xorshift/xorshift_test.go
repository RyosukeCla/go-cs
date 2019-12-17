package xorshift

import (
	"testing"
)

func Test(t *testing.T) {
	rand := NewRand(123)
	for i := 0; i < 1000; i++ {
		rand.Generate()
	}
}
