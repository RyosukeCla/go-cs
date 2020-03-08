package xorshift

import (
	"testing"

	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

func Test(t *testing.T) {
	rand := New(123)
	for i := 0; i < 1000; i++ {
		r := rand.Generate()
		if 0 > r || r >= 1 {
			t.Fatal("error", r)
		}
	}
}

func TestMu(t *testing.T) {
	size := 1000000
	rand := New(123)
	rands := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		rands = append(rands, rand.Generate())
	}

	sum := 0.0
	for _, r := range rands {
		sum += r
	}

	mu := sum / float64(size)

	if math.Abs(mu-0.5) > 0.001 {
		t.Fatal("error", mu)
	}
}
