package interpolation

import (
	"testing"

	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

func TestLinear(t *testing.T) {
	y := Linear(0, 10, 0.5)

	if math.Abs(y-5) > 0.001 {
		t.Fatal("error", y)
	}
}

func TestCosine(t *testing.T) {
	y0 := Cosine(2, 10, 0)
	y1 := Cosine(2, 10, 1)
	y := Cosine(2, 10, 0.5)

	if math.Abs(y0-2) > 0.001 {
		t.Fatal("error", y0)
	}

	if math.Abs(y1-10) > 0.001 {
		t.Fatal("error", y1)
	}

	if math.Abs(y-6) > 0.001 {
		t.Fatal("error", y)
	}
}
