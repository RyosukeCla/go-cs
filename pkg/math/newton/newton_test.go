package newton

import (
	"testing"

	"github.com/RyosukeCla/go-cs/pkg/math/basic"
)

func TestSolve(t *testing.T) {
	actual := Solve(&Args{
		F: func(x float64) float64 {
			return x*x - 2
		},
		Fdx: func(x float64) float64 {
			return 2 * x
		},
		N:  5,
		X0: 1,
	})

	expected := 1.414213562373095 // sqrt(2)
	if basic.Abs(actual-expected) > 0.00000002 {
		t.Fatal("Error")
	}
}
